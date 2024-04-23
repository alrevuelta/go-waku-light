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

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"github.com/waku-org/go-zerokit-rln/rln"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

const Endpoint = "wss://ws.cardona.zkevm-rpc.com"
const RlnContractAddress = "0x16abffcab50e8d1ff5c22b118be5c56f801dce54"

// Example data for RLN messages
const Message = "Hello World!"
const RlnEpoch = 1

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
					err := LocalMerkleProof(cfg, chunkSize, leafIndex)
					return err
				},
			},
			{
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "membership-file",
						Destination: &membershipFile,
					},
				},
				Name: "onchain-generate-rln-proof",
				Action: func(cCtx *cli.Context) error {
					err := OnchainGenerateRlnProof(cfg, membershipFile)
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
				},
				Name: "local-generate-rln-proof",
				Action: func(cCtx *cli.Context) error {
					err := LocalGenerateRlnProof(cfg, chunkSize, membershipFile)
					return err
				},
			},
			{
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "proof-file",
						Destination: &proofFile,
					},
				},
				Name: "verify-rln-proof",
				Action: func(cCtx *cli.Context) error {
					err := VerifyRlnProof(cfg, proofFile)
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

func LocalMerkleProof(cfg *Config, chunkSize uint64, leafIndex uint64) error {
	rlnInstance, err := SyncTree(cfg, chunkSize)
	if err != nil {
		return errors.Wrap(err, "error when syncing tree")
	}

	proof, err := rlnInstance.GetMerkleProof(rln.MembershipIndex(leafIndex))
	if err != nil {
		return errors.Wrap(err, "error when fetching merkle proof")
	}

	log.Info("Merkle elements for leaf ", leafIndex)
	for i, element := range proof.PathElements {
		// Bytes are reversed so that they match with the contract.
		// RLN uses little endian, while the contract uses big endian.
		bigInt := new(big.Int).SetBytes(reverseBytes(element[:]))
		log.Info("Local proof element [", i, "] ", bigInt)
	}

	log.Info("Full Merkle proof for leaf: ", leafIndex, " ", proof)

	return nil
}

// Generates an rln zk proof to be attached to a message, proving membership
// inclusion + respecting rate limits. It requires a valid rln membership that
// has been registered in the contract.
func LocalGenerateRlnProof(cfg *Config, chunkSize uint64, rlnFile string) error {
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

	proof, err := rlnInstance.GenerateProof([]byte(Message), *idCred, rln.MembershipIndex(membershipIndex), rln.Epoch{RlnEpoch})
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

func OnchainGenerateRlnProof(cfg *Config, rlnFile string) error {
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

	rlnInstance, err := rln.NewRLN()
	if err != nil {
		return errors.Wrap(err, "error when creating RLN instance")
	}

	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}

	exists, err := cfg.contract.MemberExists(callOpts, rln.Bytes32ToBigInt(idCred.IDCommitment))
	if err != nil {
		return errors.Wrap(err, "error when checking if membership exists")
	}

	if !exists {
		return errors.New("membership does not exist in the contract")
	}

	membershipIndex, err := cfg.contract.Members(callOpts, rln.Bytes32ToBigInt(idCred.IDCommitment))
	if err != nil {
		return errors.Wrap(err, "error when fetching membership index")
	}
	log.Info("Membership index found in the contract: ", membershipIndex, " for the provided commitment")

	merkleProof, err := OnchainMerkleProof(cfg, membershipIndex.Uint64())
	if err != nil {
		return errors.Wrap(err, "error when fetching merkle proof")
	}

	rlnWitness := rln.CreateWitness(
		idCred.IDSecretHash,
		[]byte(Message),
		rln.ToEpoch(RlnEpoch),
		*merkleProof)

	proof, err := rlnInstance.GenerateRLNProofWithWitness(rlnWitness)
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

func VerifyRlnProof(cfg *Config, proofFile string) error {
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

	verified, err := rlnInstance.Verify([]byte(Message), *proof, rln.BigIntToBytes32(onchainRoot))
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
