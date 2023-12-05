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
)

// some public endopoints

// drpc
// https://polygon-zkevm-testnet.drpc.org
// wss://polygon-zkevm-testnet.drpc.org

// blastapi
// https://polygon-zkevm-testnet.public.blastapi.io
// wss://polygon-zkevm-testnet.public.blastapi.io

// polygon team
// http://rpc.public.zkevm-test.net
// wss://ws.public.zkevm-test.net

const Endpoint = "wss://ws.public.zkevm-test.net"
const RlnContractAddress = "0x16aBFfCAB50E8D1ff5c22b118Be5c56F801Dce54"

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
	// ./main onchain-root
	// ./main sync-tree --chunk-size=500
	// ./main generate-proof --membership-file=membership_xxx.json
	// ./main verify-proof --proof-file=proof_xxx.json
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
					Register(cfg, privKey, amountRegister)
					return nil
				},
			},
			{
				Name: "listen",
				Action: func(cCtx *cli.Context) error {
					Listen(cfg)
					return nil
				},
			},
			{
				Name: "onchain-root",
				Action: func(cCtx *cli.Context) error {
					OnchainRoot(cfg)
					return nil
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

				Name: "sync-tree",
				Action: func(cCtx *cli.Context) error {
					SyncTree(cfg, chunkSize)
					return nil
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
				Name: "generate-proof",
				Action: func(cCtx *cli.Context) error {
					GenerateProof(cfg, chunkSize, membershipFile)
					return nil
				},
			},
			{
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "proof-file",
						Destination: &proofFile,
					},
				},
				Name: "verify-proof",
				Action: func(cCtx *cli.Context) error {
					VerifyProof(cfg, proofFile)
					return nil
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
func Register(cfg *Config, privKey string, amount int) {
	log.Info("Configured contract ", RlnContractAddress, " registering ", amount, " memberships")
	rlnInstance, err := rln.NewRLN()
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	log.Info("Preparing tx from address: ", fromAddress.Hex())
	nonce, err := cfg.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	chaindId, err := cfg.client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chaindId)
	if err != nil {
		log.Fatal(err)
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
			log.Fatal(err)
		}

		mBig := rln.Bytes32ToBigInt(m.IDCommitment)

		// Create a tx calling the update rewards root function with the new merkle root
		tx, err := cfg.contract.Register(auth, mBig)
		if err != nil {
			log.Fatal(err)
		}

		log.Info("Tx sent. Nonce: ", auth.Nonce, " Commitment: ", mBig, " TxHash: ", tx.Hash().Hex())

		rankingsJson, err := json.Marshal(m)
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile(fmt.Sprintf("membership_%s.json", mBig.String()), rankingsJson, 0644)
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(4 * time.Second)
	}
}

// Listens for new registrations and logs new root. Note that slashings are not
// monitored.
func Listen(cfg *Config) {
	log.Info("Configured contract ", RlnContractAddress)

	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}
	onchainRoot, err := cfg.contract.Root(callOpts)
	if err != nil {
		log.Fatal(err)
	}

	numLeafs, err := cfg.contract.IdCommitmentIndex(callOpts)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("There are ", numLeafs, " leafs and the root is ", onchainRoot)
	log.Info("Listening for new registrations...")

	currentBlock, err := cfg.client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: &currentBlock}
	channel := make(chan *contract.ContractMemberRegistered)

	sub, err := cfg.contract.WatchMemberRegistered(watchOpts, channel)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-channel:
			callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}
			onchainRoot, err := cfg.contract.Root(callOpts)
			if err != nil {
				log.Fatal(err)
			}
			numLeafs, err := cfg.contract.IdCommitmentIndex(callOpts)
			if err != nil {
				log.Fatal(err)
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
func OnchainRoot(cfg *Config) {
	log.Info("Configured contract ", RlnContractAddress)
	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}
	onchainRoot, err := cfg.contract.Root(callOpts)
	if err != nil {
		log.Fatal(err)
	}

	numLeafs, err := cfg.contract.IdCommitmentIndex(callOpts)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Onchain leafs: ", numLeafs)
	log.Info("Onchain root: ", onchainRoot)
}

// Generates an rln zk proof to be attached to a message, proving membership
// inclusion + respecting rate limits. It requires a valid rln membership that
// has been registered in the contract.
func GenerateProof(cfg *Config, chunkSize uint64, rlnFile string) {
	idCred := &rln.IdentityCredential{}
	jsonFile, err := os.Open(rlnFile)
	if err != nil {
		log.Fatal(err)
	}

	bb, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bb, &idCred)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Loaded commitment: ", rln.Bytes32ToBigInt(idCred.IDCommitment))

	rlnInstance := SyncTree(cfg, chunkSize)

	found := false
	membershipIndex := uint(0)
	for leafIdx := uint(0); leafIdx < rlnInstance.LeavesSet(); leafIdx++ {

		leaf, err := rlnInstance.GetLeaf(leafIdx)
		if err != nil {
			log.Fatal(err)
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
		log.Fatal("Could not find leaf in tree")
	}

	proof, err := rlnInstance.GenerateProof([]byte(Message), *idCred, rln.MembershipIndex(membershipIndex), rln.Epoch{RlnEpoch})
	if err != nil {
		log.Fatal(err)
	}

	proofJson, err := json.Marshal(proof)
	if err != nil {
		log.Fatal("error when marshalinng proof: ", err)
	}

	// Just a hash of the proof, used as filename
	hash := sha256.Sum256(proofJson)

	fileName := fmt.Sprintf("proof_%s.json", hex.EncodeToString(hash[:]))
	err = ioutil.WriteFile(fileName, proofJson, 0644)
	if err != nil {
		log.Fatal("error when writing to file: ", err)
	}

	log.Info("Proof generated succesfully and stored as ", fileName)
}

func VerifyProof(cfg *Config, proofFile string) {
	proof := &rln.RateLimitProof{}
	jsonFile, err := os.Open(proofFile)
	if err != nil {
		log.Fatal(err)
	}

	bb, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bb, &proof)
	if err != nil {
		log.Fatal(err)
	}

	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}
	onchainRoot, err := cfg.contract.Root(callOpts)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Onchain root: ", onchainRoot)

	rlnInstance, err := rln.NewRLN()
	if err != nil {
		log.Fatal(err)
	}

	verified, err := rlnInstance.Verify([]byte(Message), *proof, rln.BigIntToBytes32(onchainRoot))
	if err != nil {
		log.Fatal(err)
	}

	// When does it fail to verify?
	// * If the data (message) is different
	// * If the membership is not part of the tree
	if !verified {
		log.Fatal("Proof not verified")
	}

	metadata, err := rlnInstance.ExtractMetadata(*proof)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Proof verified succesfully")

	_ = metadata
}

// Creates an rln instance and syncs it with the contract leafs (aka memberships). This
// creates a local tree that can be used to generate proofs (required for sending messages).
func SyncTree(cfg *Config, chunkSize uint64) *rln.RLN {
	rlnInstance, err := rln.NewRLN()
	if err != nil {
		log.Fatal(err)
	}

	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}

	numLeafs, err := cfg.contract.IdCommitmentIndex(callOpts)
	if err != nil {
		log.Fatal(err)
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
			log.Fatal(err)
		}

		for _, leaf := range leafs {
			err := rlnInstance.InsertMember(rln.BigIntToBytes32(leaf))
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	log.Info("Local leafs:", rlnInstance.LeavesSet())

	myRoot, err := rlnInstance.GetMerkleRoot()
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Local root:", rln.Bytes32ToBigInt(myRoot))

	return rlnInstance
}
