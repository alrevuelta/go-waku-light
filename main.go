package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"main/contract"
	"math/big"
	"os"
	"time"

	"github.com/libp2p/go-libp2p/core/peer"
	libp2pprot "github.com/libp2p/go-libp2p/core/protocol"
	"github.com/multiformats/go-multiaddr"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"github.com/waku-org/go-zerokit-rln/rln"
	"google.golang.org/protobuf/proto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"github.com/waku-org/go-waku/waku/v2/node"
	"github.com/waku-org/go-waku/waku/v2/peerstore"
	"github.com/waku-org/go-waku/waku/v2/protocol/lightpush"
	"github.com/waku-org/go-waku/waku/v2/protocol/pb"
	rlnpb "github.com/waku-org/go-waku/waku/v2/protocol/rln/pb"
	"github.com/waku-org/go-waku/waku/v2/utils"
)

const Endpoint = "wss://ws.cardona.zkevm-rpc.com"
const RlnContractAddress = "0x520434D97e5eeD39a1F44C1f41A8024cB6138772"

type Config struct {
	client   *ethclient.Client
	contract *contract.Contract
}

func main() {
	// cli flags
	var chunkSize uint64
	var membershipFile string
	var proofFile string
	var amountRegister int
	var privKey string
	var leafIndex uint64
	var message string
	var contentTopic string
	var pubsubTopic string
	var clusterId int
	var lightpushPeer string

	executionClient, err := ethclient.Dial(Endpoint)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(RlnContractAddress)
	contract, err := contract.NewContract(address, executionClient)
	if err != nil {
		log.Fatal(err)
	}
	cfg := &Config{
		client:   executionClient,
		contract: contract,
	}

	// Examples of usage:
	// ./main register --priv-key=REPLACE_YOUR_PRIV_KEY
	// ./main listen

	// Fetched direcly from the contract
	// ./main onchain-root
	// ./main onchain-merkle-proof --leaf-index=1

	// Syncronized via events locally
	// ./main local-root --chunk-size=500
	// ./main local-merkle-proof --chunk-size=500 --leaf-index=1

	// RLN Related
	// ./main onchain-generate-rln-proof --membership-file=membership_xxx.json
	// ./main local-generate-rln-proof --membership-file=membership_xxx.json --chunk-size=500
	// ./main verify-rln-proof --proof-file=proof_xxx.json

	// Publish message via lightpush
	// ./main send-message --membership-file=membership_xxx.json --message="light client sending a rln message"
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:        "amount",
						Value:       1,
						Destination: &amountRegister,
					},
					&cli.StringFlag{
						Name:        "priv-key",
						Destination: &privKey,
					},
				},
				Name: "register",
				Action: func(cCtx *cli.Context) error {
					err := Register(cfg, privKey, amountRegister)
					return err
				},
			},
			{
				Name: "listen",
				Action: func(cCtx *cli.Context) error {
					err := Listen(cfg)
					return err
				},
			},
			{
				Name: "onchain-root",
				Action: func(cCtx *cli.Context) error {
					err := OnchainRoot(cfg)
					return err
				},
			},
			{
				Flags: []cli.Flag{
					&cli.Uint64Flag{
						Name:        "leaf-index",
						Destination: &leafIndex,
					},
				},

				Name: "onchain-merkle-proof",
				Action: func(cCtx *cli.Context) error {
					_, err := OnchainMerkleProof(cfg, leafIndex)
					return err
				},
			},
			{
				Flags: []cli.Flag{
					&cli.Uint64Flag{
						Name:        "chunk-size",
						Value:       500,
						Destination: &chunkSize,
					},
				},

				Name: "local-root",
				Action: func(cCtx *cli.Context) error {
					_, err := SyncTree(cfg, chunkSize)
					return err
				},
			},
			{
				Flags: []cli.Flag{
					&cli.Uint64Flag{
						Name:        "chunk-size",
						Value:       500,
						Destination: &chunkSize,
					},
					&cli.Uint64Flag{
						Name:        "leaf-index",
						Destination: &leafIndex,
					},
				},

				Name: "local-merkle-proof",
				Action: func(cCtx *cli.Context) error {
					_, err := LocalMerkleProof(cfg, chunkSize, leafIndex)
					return err
				},
			},
			{
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "membership-file",
						Destination: &membershipFile,
					},
					&cli.StringFlag{
						Name:        "message",
						Destination: &message,
						DefaultText: "Hello World!",
					},
					&cli.StringFlag{
						Name:        "content-topic",
						Destination: &contentTopic,
						Value:       "/basic2/1/test/proto",
					},
				},
				Name: "onchain-generate-rln-proof",
				Action: func(cCtx *cli.Context) error {
					_, err := OnchainGenerateRlnProof(cfg, membershipFile, message, contentTopic)
					return err
				},
			},
			{
				Flags: []cli.Flag{
					&cli.Uint64Flag{
						Name:        "chunk-size",
						Value:       500,
						Destination: &chunkSize,
					},
					&cli.StringFlag{
						Name:        "membership-file",
						Destination: &membershipFile,
					},
					&cli.StringFlag{
						Name:        "message",
						Destination: &message,
						DefaultText: "Hello World!",
					},
					&cli.StringFlag{
						Name:        "content-topic",
						Destination: &contentTopic,
						Value:       "/basic2/1/test/proto",
					},
				},
				Name: "local-generate-rln-proof",
				Action: func(cCtx *cli.Context) error {
					err := LocalGenerateRlnProof(cfg, chunkSize, membershipFile, message, contentTopic)
					return err
				},
			},
			{
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "proof-file",
						Destination: &proofFile,
					},
					&cli.StringFlag{
						Name:        "message",
						Destination: &message,
						DefaultText: "Hello World!",
					},
					&cli.StringFlag{
						Name:        "content-topic",
						Destination: &contentTopic,
						Value:       "/basic2/1/test/proto",
					},
				},
				Name: "verify-rln-proof",
				Action: func(cCtx *cli.Context) error {
					err := VerifyRlnProof(cfg, proofFile, message, contentTopic)
					return err
				},
			},

			{
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "membership-file",
						Destination: &membershipFile,
					},
					&cli.StringFlag{
						Name:        "message",
						Destination: &message,
					},
					&cli.StringFlag{
						Name:        "content-topic",
						Destination: &contentTopic,
						Value:       "/basic2/1/test/proto",
					},
					&cli.StringFlag{
						Name:        "lightpush-peer",
						Destination: &lightpushPeer,
						Value:       "/ip4/127.0.0.1/tcp/60000/p2p/16Uiu2HAkxTGJRgkCxgMDH4A4QBvw3q462BRkVJaPF5KQWkc1t4cp",
					},
					&cli.StringFlag{
						Name:        "pubsub-topic",
						Destination: &pubsubTopic,
						Value:       "/waku/2/rs/100/0",
					},
					&cli.IntFlag{
						Name:        "cluster-id",
						Destination: &clusterId,
						Value:       100,
					},
				},
				Name: "send-message",
				Action: func(cCtx *cli.Context) error {
					err := SendMessage(
						cfg,
						membershipFile,
						message,
						contentTopic,
						uint16(clusterId),
						lightpushPeer,
						pubsubTopic)
					return err
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// Register a new membership into the rln contract, and stores its in a json file. Note that the json
// is not a keystore, but just a custom serialized struct. Registering requires providing a valid
// account with enough funds.
func Register(cfg *Config, privKey string, amount int) error {
	log.Info("Configured contract ", RlnContractAddress, " registering ", amount, " memberships")
	rlnInstance, err := rln.NewRLN()
	if err != nil {
		return errors.Wrap(err, "error when creating RLN instance")
	}

	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return errors.Wrap(err, "error when converting private key")
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	log.Info("Preparing tx from address: ", fromAddress.Hex())
	nonce, err := cfg.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return errors.Wrap(err, "error when fetching nonce")
	}

	chaindId, err := cfg.client.NetworkID(context.Background())
	if err != nil {
		return errors.Wrap(err, "error when fetching chain id")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chaindId)
	if err != nil {
		return errors.Wrap(err, "error when creating transactor")
	}

	var i uint64 = 0
	for i = 0; i < uint64(amount); i++ {
		auth.Nonce = big.NewInt(int64(nonce + uint64(i)))
		auth.Value = big.NewInt(0)
		auth.GasPrice = nil
		auth.GasFeeCap = nil
		auth.GasTipCap = nil
		auth.NoSend = false
		auth.Context = context.Background()
		auth.GasLimit = uint64(1_000_000)

		m, err := rlnInstance.MembershipKeyGen()
		if err != nil {
			return errors.Wrap(err, "error when generating membership")
		}

		mBig := rln.Bytes32ToBigInt(m.IDCommitment)

		// Create a tx calling the update rewards root function with the new merkle root
		tx, err := cfg.contract.Register(auth, mBig)
		if err != nil {
			return errors.Wrap(err, "error when sending tx")
		}

		log.Info("Tx sent. Nonce: ", auth.Nonce, " Commitment: ", mBig, " TxHash: ", tx.Hash().Hex())

		rankingsJson, err := json.Marshal(m)
		if err != nil {
			return errors.Wrap(err, "error when marshalling membership")
		}
		err = ioutil.WriteFile(fmt.Sprintf("membership_%s.json", mBig.String()), rankingsJson, 0644)
		if err != nil {
			return errors.Wrap(err, "error when writing membership to file")
		}

		time.Sleep(4 * time.Second)
	}

	return nil
}

// Listens for new registrations and logs new root. Note that slashings are not
// monitored.
func Listen(cfg *Config) error {
	log.Info("Configured contract ", RlnContractAddress)

	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}
	onchainRoot, err := cfg.contract.Root(callOpts)
	if err != nil {
		return errors.Wrap(err, "error when fetching root")
	}

	numLeafs, err := cfg.contract.IdCommitmentIndex(callOpts)
	if err != nil {
		return errors.Wrap(err, "error when fetching num leafs")
	}

	log.Info("There are ", numLeafs, " leafs and the root is ", onchainRoot)
	log.Info("Listening for new registrations...")

	currentBlock, err := cfg.client.BlockNumber(context.Background())
	if err != nil {
		return errors.Wrap(err, "error when fetching block number")
	}
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: &currentBlock}
	channel := make(chan *contract.ContractMemberRegistered)

	sub, err := cfg.contract.WatchMemberRegistered(watchOpts, channel)
	if err != nil {
		return errors.Wrap(err, "error when watching events")
	}

	for {
		select {
		case err := <-sub.Err():
			return errors.Wrap(err, "error when watching events")
		case vLog := <-channel:
			callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}
			onchainRoot, err := cfg.contract.Root(callOpts)
			if err != nil {
				return errors.Wrap(err, "error when fetching root")
			}
			numLeafs, err := cfg.contract.IdCommitmentIndex(callOpts)
			if err != nil {
				return errors.Wrap(err, "error when fetching num leafs")
			}
			log.WithFields(log.Fields{
				"Block":         vLog.Raw.BlockNumber,
				"NewCommitment": vLog.IdCommitment,
				"NewRoot":       onchainRoot,
				"NewNumLeafs":   numLeafs,
			}).Info("New registration detected")
		}
	}
}

// Gets the merkle root from the contract and logs it.
func OnchainRoot(cfg *Config) error {
	log.Info("Configured contract ", RlnContractAddress)
	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}
	onchainRoot, err := cfg.contract.Root(callOpts)
	if err != nil {
		return errors.Wrap(err, "error when fetching root")
	}

	numLeafs, err := cfg.contract.IdCommitmentIndex(callOpts)
	if err != nil {
		return errors.Wrap(err, "error when fetching num leafs")
	}

	log.Info("Onchain leafs: ", numLeafs)
	log.Info("Onchain root: ", onchainRoot)

	return nil
}

func OnchainMerkleProof(cfg *Config, leafIndex uint64) (*rln.MerkleProof, error) {
	log.Info("Configured contract ", RlnContractAddress)
	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}

	merkleProofElements, err := cfg.contract.MerkleProofElements(callOpts, big.NewInt(0).SetUint64(leafIndex))
	if err != nil {
		return nil, errors.Wrap(err, "error when fetching merkle proof elements")
	}

	bytePathElements := make([]rln.MerkleNode, len(merkleProofElements))

	log.Info("Merkle elements for leaf ", leafIndex)
	for i, element := range merkleProofElements {
		log.Info("Raw from contract [", i, "] ", element)
		bytePathElements[i] = rln.BigIntToBytes32(element)
	}

	pathIndexes := make([]uint8, len(merkleProofElements))
	for i := 0; i < len(merkleProofElements); i++ {
		index := (leafIndex >> i) & 1
		pathIndexes[i] = uint8(index)
	}

	merkleProof := &rln.MerkleProof{
		PathElements: bytePathElements,
		PathIndexes:  pathIndexes,
	}

	log.Info("Full Merkle proof for leaf ", leafIndex, " ", merkleProof)

	return merkleProof, nil
}

func LocalMerkleProof(cfg *Config, chunkSize uint64, leafIndex uint64) (*rln.MerkleProof, error) {
	rlnInstance, err := SyncTree(cfg, chunkSize)
	if err != nil {
		return nil, errors.Wrap(err, "error when syncing tree")
	}

	proof, err := rlnInstance.GetMerkleProof(rln.MembershipIndex(leafIndex))
	if err != nil {
		return nil, errors.Wrap(err, "error when fetching merkle proof")
	}

	log.Info("Merkle elements for leaf ", leafIndex)
	for i, element := range proof.PathElements {
		// Bytes are reversed so that they match with the contract.
		// RLN uses little endian, while the contract uses big endian.
		bigInt := new(big.Int).SetBytes(reverseBytes(element[:]))
		log.Info("Local proof element [", i, "] ", bigInt)
	}

	log.Info("Full Merkle proof for leaf: ", leafIndex, " ", proof)

	return &proof, nil
}

// Generates an rln zk proof to be attached to a message, proving membership
// inclusion + respecting rate limits. It requires a valid rln membership that
// has been registered in the contract.
func LocalGenerateRlnProof(
	cfg *Config,
	chunkSize uint64,
	rlnFile string,
	message string,
	contentTopic string) error {

	idCred := &rln.IdentityCredential{}
	jsonFile, err := os.Open(rlnFile)
	if err != nil {
		return errors.Wrap(err, "error when opening file")
	}

	bb, err := io.ReadAll(jsonFile)
	if err != nil {
		return errors.Wrap(err, "error when reading file")
	}
	err = json.Unmarshal(bb, &idCred)
	if err != nil {
		return errors.Wrap(err, "error when unmarshalling file")
	}

	log.Info("Loaded commitment: ", rln.Bytes32ToBigInt(idCred.IDCommitment))

	rlnInstance, err := SyncTree(cfg, chunkSize)
	if err != nil {
		return errors.Wrap(err, "error when syncing tree")
	}

	membershipIndex, err := findMembershipInTree(rlnInstance, idCred)
	if err != nil {
		return errors.Wrap(err, "error when finding membership in tree")
	}

	// Topic and message are used as inputs
	// https://github.com/waku-org/go-waku/blob/v0.9.0/waku/v2/protocol/rln/common.go#L33-L40
	x := append([]byte(message), []byte(contentTopic)...)

	proof, err := rlnInstance.GenerateProof(x, *idCred, rln.MembershipIndex(membershipIndex), rln.GetCurrentEpoch())
	if err != nil {
		return errors.Wrap(err, "error when generating proof")
	}

	proofJson, err := json.Marshal(proof)
	if err != nil {
		return errors.Wrap(err, "error when marshalling proof")
	}

	// Just a hash of the proof, used as filename
	hash := sha256.Sum256(proofJson)

	fileName := fmt.Sprintf("proof_%s.json", hex.EncodeToString(hash[:]))
	err = ioutil.WriteFile(fileName, proofJson, 0644)
	if err != nil {
		return errors.Wrap(err, "error when writing to file")
	}

	log.Info("Proof generated succesfully and stored as ", fileName)

	return nil
}

func OnchainGenerateRlnProof(
	cfg *Config,
	membershipFile string,
	message string,
	contentTopic string) (*rln.RateLimitProof, error) {

	idCred := &rln.IdentityCredential{}
	jsonFile, err := os.Open(membershipFile)
	if err != nil {
		return nil, errors.Wrap(err, "error when opening file")
	}

	bb, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, errors.Wrap(err, "error when reading file")
	}
	err = json.Unmarshal(bb, &idCred)
	if err != nil {
		return nil, errors.Wrap(err, "error when unmarshalling file")
	}

	rlnInstance, err := rln.NewRLN()
	if err != nil {
		return nil, errors.Wrap(err, "error when creating RLN instance")
	}

	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}

	exists, err := cfg.contract.MemberExists(callOpts, rln.Bytes32ToBigInt(idCred.IDCommitment))
	if err != nil {
		return nil, errors.Wrap(err, "error when checking if membership exists")
	}

	if !exists {
		return nil, errors.New("membership does not exist in the contract")
	}

	membershipIndex, err := cfg.contract.Members(callOpts, rln.Bytes32ToBigInt(idCred.IDCommitment))
	if err != nil {
		return nil, errors.Wrap(err, "error when fetching membership index")
	}
	log.Info("Membership index found in the contract: ", membershipIndex, " for the provided commitment")

	merkleProof, err := OnchainMerkleProof(cfg, membershipIndex.Uint64())
	if err != nil {
		return nil, errors.Wrap(err, "error when fetching merkle proof")
	}

	// Topic and message are used as inputs
	// https://github.com/waku-org/go-waku/blob/v0.9.0/waku/v2/protocol/rln/common.go#L33-L40
	x := append([]byte(message), []byte(contentTopic)...)

	rlnWitness := rln.CreateWitness(
		idCred.IDSecretHash,
		x,
		rln.GetCurrentEpoch(),
		*merkleProof)

	proof, err := rlnInstance.GenerateRLNProofWithWitness(rlnWitness)
	if err != nil {
		return nil, errors.Wrap(err, "error when generating proof")
	}

	proofJson, err := json.Marshal(proof)
	if err != nil {
		return nil, errors.Wrap(err, "error when marshalling proof")
	}

	// Just a hash of the proof, used as filename
	hash := sha256.Sum256(proofJson)

	fileName := fmt.Sprintf("proof_%s.json", hex.EncodeToString(hash[:]))
	err = ioutil.WriteFile(fileName, proofJson, 0644)
	if err != nil {
		return nil, errors.Wrap(err, "error when writing to file")
	}

	log.Info("Proof generated succesfully and stored as ", fileName)

	return proof, nil
}

func VerifyRlnProof(cfg *Config, proofFile string, message string, contentTopic string) error {
	proof := &rln.RateLimitProof{}
	jsonFile, err := os.Open(proofFile)
	if err != nil {
		return errors.Wrap(err, "error when opening file")
	}

	bb, err := io.ReadAll(jsonFile)
	if err != nil {
		return errors.Wrap(err, "error when reading file")
	}
	err = json.Unmarshal(bb, &proof)
	if err != nil {
		return errors.Wrap(err, "error when unmarshalling file")
	}

	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}
	onchainRoot, err := cfg.contract.Root(callOpts)
	if err != nil {
		return errors.Wrap(err, "error when fetching root")
	}

	log.Info("Onchain root: ", onchainRoot)

	rlnInstance, err := rln.NewRLN()
	if err != nil {
		return errors.Wrap(err, "error when creating RLN instance")
	}

	// Topic and message are used as inputs
	// https://github.com/waku-org/go-waku/blob/v0.9.0/waku/v2/protocol/rln/common.go#L33-L40
	x := append([]byte(message), []byte(contentTopic)...)

	verified, err := rlnInstance.Verify(x, *proof, rln.BigIntToBytes32(onchainRoot))
	if err != nil {
		return errors.Wrap(err, "error when verifying proof")
	}

	// When does it fail to verify?
	// * If the data (message) is different
	// * If the membership is not part of the tree
	if !verified {
		return errors.New("proof verification failed")
	}

	metadata, err := rlnInstance.ExtractMetadata(*proof)
	if err != nil {
		return errors.Wrap(err, "error when extracting metadata")
	}
	_ = metadata

	log.Info("Proof verified succesfully")

	return nil
}

// Creates an rln instance and syncs it with the contract leafs (aka memberships). This
// creates a local tree that can be used to generate proofs (required for sending messages).
func SyncTree(cfg *Config, chunkSize uint64) (*rln.RLN, error) {
	rlnInstance, err := rln.NewRLN()
	if err != nil {
		return nil, errors.Wrap(err, "error when creating RLN instance")
	}

	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}

	numLeafs, err := cfg.contract.IdCommitmentIndex(callOpts)
	if err != nil {
		return nil, errors.Wrap(err, "error when fetching num leafs")
	}

	for i := uint64(0); i < numLeafs.Uint64(); i += chunkSize {
		start := big.NewInt(0).SetUint64(i)
		end := big.NewInt(0).SetUint64(i + chunkSize)

		if end.Cmp(numLeafs) > 0 {
			end.Set(numLeafs)
		}
		log.Info("Fetching from ", start, " to ", end, " out of ", numLeafs, " leafs")
		leafs, err := cfg.contract.GetCommitments(callOpts, start, end)
		if err != nil {
			return nil, errors.Wrap(err, "error when fetching commitments")
		}

		for _, leaf := range leafs {
			err := rlnInstance.InsertMember(rln.BigIntToBytes32(leaf))
			if err != nil {
				return nil, errors.Wrap(err, "error when inserting member")
			}
		}
	}

	log.Info("Local leafs: ", rlnInstance.LeavesSet())

	myRoot, err := rlnInstance.GetMerkleRoot()
	if err != nil {
		return nil, errors.Wrap(err, "error when fetching merkle root")
	}

	log.Info("Local root: ", rln.Bytes32ToBigInt(myRoot))

	return rlnInstance, nil
}

func findMembershipInTree(rlnInstance *rln.RLN, idCred *rln.IdentityCredential) (uint, error) {
	found := false
	membershipIndex := uint(0)
	for leafIdx := uint(0); leafIdx < rlnInstance.LeavesSet(); leafIdx++ {

		leaf, err := rlnInstance.GetLeaf(leafIdx)
		if err != nil {
			return 0, errors.Wrap(err, "error when fetching leaf")
		}

		leafBig := rln.Bytes32ToBigInt(leaf)
		mineBig := rln.Bytes32ToBigInt(idCred.IDCommitment)

		if leafBig.Cmp(mineBig) == 0 {
			log.Info("Found leaf in tree: ", leafBig)
			found = true
			membershipIndex = leafIdx
			break
		}
	}

	if !found {
		return 0, errors.New("membership not found in tree")
	}

	return membershipIndex, nil
}

func reverseBytes(b []byte) []byte {
	reversed := make([]byte, len(b))
	copy(reversed, b)

	for i := len(reversed)/2 - 1; i >= 0; i-- {
		opp := len(reversed) - 1 - i
		reversed[i], reversed[opp] = reversed[opp], reversed[i]
	}

	return reversed
}

// See: https://github.com/waku-org/go-waku/blob/master/examples/basic-light-client/main.go
// Note that this requires a running waku node with lightpush enabled at localhost.
func SendMessage(
	cfg *Config,
	membershipFile string,
	message string,
	contentTopic string,
	clusterId uint16,
	lightpushPeer string,
	pubsubTopic string) error {

	wakuNode, err := node.New(node.WithClusterID(clusterId))
	if err != nil {
		return errors.Wrap(err, "error when creating waku node")
	}

	if err := wakuNode.Start(context.Background()); err != nil {
		return errors.Wrap(err, "error when starting waku node")
	}

	rlnProof, err := OnchainGenerateRlnProof(cfg, membershipFile, message, contentTopic)
	if err != nil {
		return errors.Wrap(err, "error when generating rln proof")
	}

	serializedRlnProof, err := serializeRLNProof(rlnProof)
	if err != nil {
		return errors.Wrap(err, "error when serializing rln proof")
	}

	msg := &pb.WakuMessage{
		Payload: []byte(message),
		//Version:      &uint32(0),
		ContentTopic:   contentTopic,
		Timestamp:      utils.GetUnixEpoch(),
		RateLimitProof: serializedRlnProof,
	}

	peerAddr, err := multiaddr.NewMultiaddr(lightpushPeer)
	if err != nil {
		return errors.Wrap(err, "error when creating multiaddr")
	}

	_, err = wakuNode.AddPeer(peerAddr, peerstore.Static, []string{pubsubTopic}, []libp2pprot.ID{lightpush.LightPushID_v20beta1}...)
	if err != nil {
		return errors.Wrap(err, "error when adding peer")
	}
	peerId, err := peer.AddrInfoFromP2pAddr(peerAddr)
	if err != nil {
		return errors.Wrap(err, "error when getting peer id")
	}

	log.WithFields(log.Fields{
		"Epoch":         rlnProof.Epoch,
		"Nullifier":     rlnProof.Nullifier,
		"RLNIdentifier": rlnProof.RLNIdentifier,
		"ShareX":        rlnProof.ShareX,
		"ShareY":        rlnProof.ShareY,
		"MerkleRoot":    rlnProof.MerkleRoot,
		"Proof":         rlnProof.Proof,
	}).Info("RLN Proof info")

	// Publish our message via lightpush, using our locally crafted RLN proof
	log.WithFields(log.Fields{
		"LPPeer":       peerId.ID,
		"PubsubTopic":  pubsubTopic,
		"Payload":      string(msg.Payload),
		"RLNProof":     msg.RateLimitProof,
		"ContentTopic": contentTopic,
		"Timestamp":    msg.Timestamp,
	}).Info("Publishing via lightpush")

	msgId, err := wakuNode.Lightpush().Publish(context.Background(), msg, lightpush.WithPeer(peerId.ID), lightpush.WithPubSubTopic(pubsubTopic))
	if err != nil {
		return errors.Wrap(err, "error when publishing message")
	}

	log.Info("Message sent with id: ", msgId)

	return nil
}

// A mix of:
// https://github.com/waku-org/go-waku/blob/8805f6cc45ff8c3c9d3d479d3fa8f5920fdc588f/waku/v2/protocol/rln/waku_rln_relay.go#L215-L218
// https://github.com/waku-org/go-waku/blob/8805f6cc45ff8c3c9d3d479d3fa8f5920fdc588f/waku/v2/protocol/rln/waku_rln_relay.go#L288-L301
func serializeRLNProof(proof *rln.RateLimitProof) ([]byte, error) {

	test := &rlnpb.RateLimitProof{
		Proof:         proof.Proof[:],
		MerkleRoot:    proof.MerkleRoot[:],
		Epoch:         proof.Epoch[:],
		ShareX:        proof.ShareX[:],
		ShareY:        proof.ShareY[:],
		Nullifier:     proof.Nullifier[:],
		RlnIdentifier: proof.RLNIdentifier[:],
	}

	ser, err := proto.Marshal(test)
	if err != nil {
		return nil, errors.Wrap(err, "error when marshalling proof")
	}

	return ser, nil
}
