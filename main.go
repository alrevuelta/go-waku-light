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
	"math/big"
	"os"
	"time"

	"github.com/alrevuelta/go-waku-light/contract"
	"github.com/libp2p/go-libp2p/core/peer"
	libp2pprot "github.com/libp2p/go-libp2p/core/protocol"
	"github.com/multiformats/go-multiaddr"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"google.golang.org/protobuf/proto"

	// TODO: The go-waku version that shall be imported is a custom one
	// otherwise there is a problem with go-zerokit since go-waku-light
	// uses a version that go-waku does not use. Temporal hack
	// TODO: At some point remove the dependancy from go-waku
	"github.com/waku-org/go-waku/waku/v2/node"
	"github.com/waku-org/go-waku/waku/v2/peerstore"
	"github.com/waku-org/go-waku/waku/v2/protocol/lightpush"
	"github.com/waku-org/go-waku/waku/v2/protocol/pb"
	rlnpb "github.com/waku-org/go-waku/waku/v2/protocol/rln/pb"
	"github.com/waku-org/go-waku/waku/v2/utils"
	"github.com/waku-org/go-zerokit-rln/rln"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

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
	var ethEndpoint string
	var contractAddress string
	var userMessageLimit uint64
	var messageEverySecs uint64
	var amountMessageToSend uint64
	var epochSizeSecs uint64

	// Examples of usage:
	// ./go-waku-light register --priv-key=REPLACE_YOUR_PRIV_KEY --user-message-limit=5
	// ./go-waku-light listen

	// Fetched direcly from the contract
	// ./go-waku-light onchain-root
	// ./go-waku-light onchain-merkle-proof --leaf-index=1

	// Syncronized via events locally
	// ./go-waku-light local-root --chunk-size=500
	// ./go-waku-light local-merkle-proof --chunk-size=500 --leaf-index=1

	// RLN Related
	// ./go-waku-light onchain-generate-rln-proof --membership-file=membership_xxx.json
	// ./go-waku-light local-generate-rln-proof --membership-file=membership_xxx.json --chunk-size=500
	// ./go-waku-light verify-rln-proof --proof-file=proof_xxx.json

	// Publish multiple messages in an infinite loop registering also the membership
	// Use --amount-message-to-send=0 to send messages indefinitely
	// ./go-waku-light send-messages-loop \
	// --priv-key=SOME_PRIV_KEY \
	// --user-message-limit=5 \
	// --message="light client sending a rln message" \
	// --content-topic=/basic2/1/test/proto \
	// --pubsub-topic=/waku/2/rs/100/0 \
	// --cluster-id=100 \
	// --lightpush-peer=/ip4/127.0.0.1/tcp/60000/p2p/16Uiu2HAkxTGJRgkCxgMDH4A4QBvw3q462BRkVJaPF5KQWkc1t4cp \
	// --message-every-secs=5 \
	// --epoch-size-secs=60
	app := &cli.App{
		Before: func(c *cli.Context) error {
			log.Info("Arguments:")
			for _, arg := range c.Args().Slice() {
				log.Info(arg)
			}
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "eth-endpoint",
				Value:       "wss://ws.cardona.zkevm-rpc.com",
				Destination: &ethEndpoint,
			},
			&cli.StringFlag{
				Name:        "contract-address",
				Value:       "0x1ae47AAb605E3639cA88ce8F6183C3035Eb60c62",
				Destination: &contractAddress,
			},
		},
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
						Required:    true,
					},
					&cli.Uint64Flag{
						Name:        "user-message-limit",
						Value:       5,
						Destination: &userMessageLimit,
					},
				},
				Name: "register",
				Action: func(cCtx *cli.Context) error {
					cfg, err := CreateConfig(ethEndpoint, contractAddress)
					if err != nil {
						return errors.Wrap(err, "error when creating config")
					}
					_, err = Register(cfg, privKey, amountRegister, uint32(userMessageLimit))
					return err
				},
			},
			{
				Name: "listen",
				Action: func(cCtx *cli.Context) error {
					cfg, err := CreateConfig(ethEndpoint, contractAddress)
					if err != nil {
						return errors.Wrap(err, "error when creating config")
					}
					err = Listen(cfg)
					return err
				},
			},
			{
				Name: "onchain-root",
				Action: func(cCtx *cli.Context) error {
					cfg, err := CreateConfig(ethEndpoint, contractAddress)
					if err != nil {
						return errors.Wrap(err, "error when creating config")
					}
					err = OnchainRoot(cfg)
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
					cfg, err := CreateConfig(ethEndpoint, contractAddress)
					if err != nil {
						return errors.Wrap(err, "error when creating config")
					}
					_, err = OnchainMerkleProof(cfg, leafIndex)
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
					cfg, err := CreateConfig(ethEndpoint, contractAddress)
					if err != nil {
						return errors.Wrap(err, "error when creating config")
					}
					_, err = SyncTree(cfg, chunkSize)
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
					cfg, err := CreateConfig(ethEndpoint, contractAddress)
					if err != nil {
						return errors.Wrap(err, "error when creating config")
					}
					_, err = LocalMerkleProof(cfg, chunkSize, leafIndex)
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
					&cli.Uint64Flag{
						Name:        "epoch-size-secs",
						Destination: &epochSizeSecs,
						Value:       1,
					},
				},
				Name: "onchain-generate-rln-proof",
				Action: func(cCtx *cli.Context) error {
					cfg, err := CreateConfig(ethEndpoint, contractAddress)
					if err != nil {
						return errors.Wrap(err, "error when creating config")
					}
					messageId := uint32(0)
					_, _, err = OnchainGenerateRlnProof(cfg, membershipFile, message, contentTopic, messageId, epochSizeSecs)
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
					&cli.Uint64Flag{
						Name:        "epoch-size-secs",
						Destination: &epochSizeSecs,
						Value:       1,
					},
				},
				Name: "local-generate-rln-proof",
				Action: func(cCtx *cli.Context) error {
					cfg, err := CreateConfig(ethEndpoint, contractAddress)
					if err != nil {
						return errors.Wrap(err, "error when creating config")
					}
					err = LocalGenerateRlnProof(cfg, chunkSize, membershipFile, message, contentTopic, epochSizeSecs)
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
					&cli.Uint64Flag{
						Name:        "epoch-size-secs",
						Destination: &epochSizeSecs,
						Value:       1,
					},
				},
				Name: "verify-rln-proof",
				Action: func(cCtx *cli.Context) error {
					cfg, err := CreateConfig(ethEndpoint, contractAddress)
					if err != nil {
						return errors.Wrap(err, "error when creating config")
					}
					err = VerifyRlnProof(cfg, proofFile, message, contentTopic)
					return err
				},
			},
			{
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "priv-key",
						Destination: &privKey,
						Required:    true,
					},
					&cli.Uint64Flag{
						Name:        "user-message-limit",
						Value:       5,
						Destination: &userMessageLimit,
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
						Name:        "pubsub-topic",
						Destination: &pubsubTopic,
						Value:       "/waku/2/rs/100/0",
					},
					&cli.IntFlag{
						Name:        "cluster-id",
						Destination: &clusterId,
						Value:       100,
					},
					&cli.StringFlag{
						Name:        "lightpush-peer",
						Destination: &lightpushPeer,
						Value:       "/ip4/127.0.0.1/tcp/60000/p2p/16Uiu2HAkxTGJRgkCxgMDH4A4QBvw3q462BRkVJaPF5KQWkc1t4cp",
					},
					&cli.Uint64Flag{
						Name:        "message-every-secs",
						Value:       10,
						Destination: &messageEverySecs,
					},
					&cli.Uint64Flag{
						Name:        "amount-message-to-send",
						Value:       0,
						Destination: &amountMessageToSend,
					},
					&cli.Uint64Flag{
						Name:        "epoch-size-secs",
						Destination: &epochSizeSecs,
						Value:       1,
					},
				},
				Name: "send-messages-loop",
				Action: func(cCtx *cli.Context) error {
					cfg, err := CreateConfig(ethEndpoint, contractAddress)
					if err != nil {
						return errors.Wrap(err, "error when creating config")
					}

					membershipFiles, err := Register(cfg, privKey, 1, uint32(userMessageLimit))
					if err != nil {
						return errors.Wrap(err, "error when registering memberships")
					}

					log.Info("Membership registered: ", membershipFiles)

					wakuNode, peerId, err := CreateLightClient(uint16(clusterId), lightpushPeer, pubsubTopic)
					if err != nil {
						return errors.Wrap(err, "error when creating light client")
					}

					msgId := uint32(0)
					numSent := uint64(0)
					for {
						// TODO: Allow to send depending on size
						messageToSend := fmt.Sprintf("%s: %d", message, msgId)

						err = SendMessage(
							cfg,
							membershipFiles[0],
							messageToSend,
							contentTopic,
							peerId,
							pubsubTopic,
							msgId,
							wakuNode,
							epochSizeSecs)

						if err != nil {
							return errors.Wrap(err, "error when sending message")
						}

						msgId++
						numSent++

						// If we sent the amount of messages requested, we stop
						if amountMessageToSend == numSent {
							log.Info("Sent all messages, completed. Exiting...")
							return nil
						}

						if msgId == uint32(userMessageLimit) {
							msgId = 0
						}

						log.Info("Message sent: ", messageToSend, " sleeping ", messageEverySecs, " seconds...")
						time.Sleep(time.Duration(messageEverySecs) * time.Second)
					}
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func CreateConfig(endpoint string, contractAddress string) (*Config, error) {
	log.Info("Connecting to ", endpoint, " and contract ", contractAddress)
	executionClient, err := ethclient.Dial(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "error when creating client")
	}

	address := common.HexToAddress(contractAddress)
	contract, err := contract.NewContract(address, executionClient)
	if err != nil {
		return nil, errors.Wrap(err, "error when creating contract")
	}
	cfg := &Config{
		client:   executionClient,
		contract: contract,
	}

	return cfg, nil
}

// Register a new membership into the rln contract, and stores its in a json file. Note that the json
// is not a keystore, but just a custom serialized struct. Registering requires providing a valid
// account with enough funds.
func Register(cfg *Config, privKey string, amount int, userMessageLimit uint32) ([]string, error) {
	log.Info("Registering ", amount, " memberships")
	log.Info("User message limit: ", userMessageLimit)

	filesNames := make([]string, 0)

	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}
	maxMessageLimit, err := cfg.contract.MAXMESSAGELIMIT(callOpts)
	if err != nil {
		return filesNames, errors.Wrap(err, "error when fetching max message limit")
	}

	// Ensure we dont attempt to register a user with a message limit higher than the contract allows
	if userMessageLimit > maxMessageLimit {
		return filesNames, errors.New(fmt.Sprintf("user message limit is too high. Requested: %d, Max: %d",
			userMessageLimit, maxMessageLimit))
	}

	rlnInstance, err := rln.NewRLN()
	if err != nil {
		return filesNames, errors.Wrap(err, "error when creating RLN instance")
	}

	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return filesNames, errors.Wrap(err, "error when converting private key")
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return filesNames, errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	log.Info("Preparing tx from address: ", fromAddress.Hex())
	nonce, err := cfg.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return filesNames, errors.Wrap(err, "error when fetching nonce")
	}

	chaindId, err := cfg.client.NetworkID(context.Background())
	if err != nil {
		return filesNames, errors.Wrap(err, "error when fetching chain id")
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chaindId)
	if err != nil {
		return filesNames, errors.Wrap(err, "error when creating transactor")
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
			return filesNames, errors.Wrap(err, "error when generating membership")
		}

		mBig := rln.Bytes32ToBigInt(m.IDCommitment)

		// Create a tx calling the update rewards root function with the new merkle root
		tx, err := cfg.contract.Register(auth, mBig, userMessageLimit)
		if err != nil {
			return filesNames, errors.Wrap(err, "error when sending tx")
		}

		log.Info("Tx sent. Nonce: ", auth.Nonce, " Commitment: ", mBig, " UserMessageLimit: ", userMessageLimit, " TxHash: ", tx.Hash().Hex())

		log.Info("Waiting for tx to be validated...")

		// Leave 60 seconds for the tx to be validated
		deadline := time.Now().Add(60 * time.Second)
		ctx, cancelCtx := context.WithDeadline(context.Background(), deadline)
		defer cancelCtx()

		// It stops waiting when the context is canceled.
		receipt, err := bind.WaitMined(ctx, cfg.client, tx)
		if ctx.Err() != nil {
			return filesNames, errors.Wrap(err,
				fmt.Sprint("timeout expired waiting for tx to be validated txHash: ",
					tx.Hash().Hex()))
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			return filesNames, errors.Wrap(err,
				fmt.Sprintf("tx failed err: %d hash: %s", receipt.Status, tx.Hash().Hex()))
		}

		log.Info("Tx validated. Receipt: ", receipt)

		rankingsJson, err := json.Marshal(m)
		if err != nil {
			return filesNames, errors.Wrap(err, "error when marshalling membership")
		}

		fileName := fmt.Sprintf("membership_%s.json", mBig.String())
		err = ioutil.WriteFile(fileName, rankingsJson, 0644)
		if err != nil {
			return filesNames, errors.Wrap(err, "error when writing membership to file")
		}

		filesNames = append(filesNames, fileName)
	}

	// We wait some time to ensure the root is in nwaku nodes
	log.Info("Waiting for the root to be propagated to the network...")
	time.Sleep(6 * time.Second)

	return filesNames, nil
}

// Listens for new registrations and logs new root. Note that slashings are not
// monitored.
func Listen(cfg *Config) error {
	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}
	onchainRoot, err := cfg.contract.Root(callOpts)
	if err != nil {
		return errors.Wrap(err, "error when fetching root")
	}

	numLeafs, err := cfg.contract.CommitmentIndex(callOpts)
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
			numLeafs, err := cfg.contract.CommitmentIndex(callOpts)
			if err != nil {
				return errors.Wrap(err, "error when fetching num leafs")
			}
			log.WithFields(log.Fields{
				"Block":         vLog.Raw.BlockNumber,
				"NewCommitment": vLog.RateCommitment,
				"NewRoot":       onchainRoot,
				"NewNumLeafs":   numLeafs,
			}).Info("New registration detected")
		}
	}
}

// Gets the merkle root from the contract and logs it.
func OnchainRoot(cfg *Config) error {
	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}
	onchainRoot, err := cfg.contract.Root(callOpts)
	if err != nil {
		return errors.Wrap(err, "error when fetching root")
	}

	numLeafs, err := cfg.contract.CommitmentIndex(callOpts)
	if err != nil {
		return errors.Wrap(err, "error when fetching num leafs")
	}

	log.Info("Onchain leafs: ", numLeafs)
	log.Info("Onchain root: ", onchainRoot)

	return nil
}

func OnchainMerkleProof(cfg *Config, leafIndex uint64) (*rln.MerkleProof, error) {
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
	contentTopic string,
	epochSizeSecs uint64) error {

	idCred := &rln.IdentityCredential{}
	jsonFile, err := os.Open(rlnFile)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error when opening file %s", rlnFile))
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
	data := append([]byte(message), []byte(contentTopic)...)

	// TODO: Hardcoded by now
	messageId := uint32(1)

	proof, err := rlnInstance.GenerateProof(
		data,
		*idCred,
		rln.MembershipIndex(membershipIndex),
		rln.GetCurrentEpoch(epochSizeSecs),
		messageId)

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
	contentTopic string,
	messageId uint32,
	epochSizeSecs uint64) (*rln.RateLimitProof, *rln.Epoch, error) {

	idCred := &rln.IdentityCredential{}
	jsonFile, err := os.Open(membershipFile)
	if err != nil {
		return nil, nil, errors.Wrap(err, fmt.Sprintf("error when opening file %s", membershipFile))
	}

	bb, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error when reading file")
	}
	err = json.Unmarshal(bb, &idCred)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error when unmarshalling file")
	}

	rlnInstance, err := rln.NewRLN()
	if err != nil {
		return nil, nil, errors.Wrap(err, "error when creating RLN instance")
	}

	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}

	exists, err := cfg.contract.MemberExists(callOpts, rln.Bytes32ToBigInt(idCred.IDCommitment))
	if err != nil {
		return nil, nil, errors.Wrap(err, "error when checking if membership exists")
	}

	if !exists {
		return nil, nil, errors.New("membership does not exist in the contract")
	}

	userMessageLimit, membershipIndex, rateCommitment, err := cfg.contract.IdCommitmentToMetadata(callOpts, rln.Bytes32ToBigInt(idCred.IDCommitment))
	if err != nil {
		return nil, nil, errors.Wrap(err, "error when fetching membership index")
	}

	log.WithFields(log.Fields{
		"UserMessageLimit": userMessageLimit,
		"MembershipIndex":  membershipIndex,
		"RateCommitment":   rateCommitment,
	}).Info("Commitment metadata")

	merkleProof, err := OnchainMerkleProof(cfg, uint64(membershipIndex))
	if err != nil {
		return nil, nil, errors.Wrap(err, "error when fetching merkle proof")
	}

	// Topic and message are used as inputs
	// https://github.com/waku-org/go-waku/blob/v0.9.0/waku/v2/protocol/rln/common.go#L33-L40
	data := append([]byte(message), []byte(contentTopic)...)

	epoch := rln.GetCurrentEpoch(epochSizeSecs)

	rlnWitness, err := rlnInstance.CreateWitness(
		idCred.IDSecretHash,
		userMessageLimit,
		messageId,
		data,
		epoch,
		*merkleProof)

	if err != nil {
		return nil, nil, errors.Wrap(err, "error when creating witness")
	}

	proof, err := rlnInstance.GenerateRLNProofWithWitness(rlnWitness)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error when generating proof")
	}

	proofJson, err := json.Marshal(proof)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error when marshalling proof")
	}

	// Just a hash of the proof, used as filename
	hash := sha256.Sum256(proofJson)

	fileName := fmt.Sprintf("proof_%s.json", hex.EncodeToString(hash[:]))
	err = ioutil.WriteFile(fileName, proofJson, 0644)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error when writing to file")
	}

	log.Info("Proof generated succesfully and stored as ", fileName)

	return proof, &epoch, nil
}

func VerifyRlnProof(cfg *Config, proofFile string, message string, contentTopic string) error {
	proof := &rln.RateLimitProof{}
	jsonFile, err := os.Open(proofFile)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("error when opening file %s", proofFile))
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

	numLeafs, err := cfg.contract.CommitmentIndex(callOpts)
	if err != nil {
		return nil, errors.Wrap(err, "error when fetching num leafs")
	}

	for i := uint64(0); i < uint64(numLeafs-1); i += chunkSize {
		start := i
		end := i + chunkSize

		if end > uint64(numLeafs-1) {
			end = uint64(numLeafs - 1)
		}
		log.Info("Fetching memberships [", start, ",", end, "] out of ", numLeafs, " leafs")

		leafs, err := cfg.contract.GetCommitments(callOpts, uint32(start), uint32(end))
		if err != nil {
			return nil, errors.Wrap(err, "error when fetching commitments")
		}

		for _, leaf := range leafs {
			err := rlnInstance.InsertRawLeaf(rln.BigIntToBytes32(leaf))
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

func CreateLightClient(clusterId uint16, lightpushPeer string, pubsubTopic string) (*node.WakuNode, *peer.AddrInfo, error) {
	wakuNode, err := node.New(node.WithClusterID(clusterId))
	if err != nil {
		return nil, nil, errors.Wrap(err, "error when creating waku node")
	}

	if err := wakuNode.Start(context.Background()); err != nil {
		return nil, nil, errors.Wrap(err, "error when starting waku node")
	}

	peerAddr, err := multiaddr.NewMultiaddr(lightpushPeer)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error when creating multiaddr")
	}

	_, err = wakuNode.AddPeer(peerAddr, peerstore.Static, []string{pubsubTopic}, []libp2pprot.ID{lightpush.LightPushID_v20beta1}...)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error when adding peer")
	}
	peerId, err := peer.AddrInfoFromP2pAddr(peerAddr)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error when getting peer id")
	}

	return wakuNode, peerId, nil
}

// See: https://github.com/waku-org/go-waku/blob/master/examples/basic-light-client/main.go
// Note that this requires a running waku node with lightpush enabled at localhost.
func SendMessage(
	cfg *Config,
	membershipFile string,
	message string,
	contentTopic string,
	peerId *peer.AddrInfo,
	pubsubTopic string,
	messageId uint32,
	wakuNode *node.WakuNode,
	epochSizeSecs uint64) error {

	log.Info("Selected membership file: ", membershipFile)

	rlnProof, epoch, err := OnchainGenerateRlnProof(cfg, membershipFile, message, contentTopic, messageId, epochSizeSecs)
	if err != nil {
		return errors.Wrap(err, "error when generating rln proof")
	}

	serializedRlnProof, err := serializeRLNProof(rlnProof, *epoch, rln.RLN_IDENTIFIER[:])
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

	log.WithFields(log.Fields{
		"Proof":             rlnProof.Proof,
		"MerkleRoot":        rlnProof.MerkleRoot,
		"ExternalNullifier": rlnProof.ExternalNullifier,
		"ShareX":            rlnProof.ShareX,
		"ShareY":            rlnProof.ShareY,
		"Nullifier":         rlnProof.Nullifier,
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

// https://github.com/waku-org/nwaku/blob/v0.28.0/waku/waku_rln_relay/protocol_types.nim#L35
// serialized as: https://github.com/waku-org/nwaku/blob/v0.28.0/waku/waku_rln_relay/protocol_types.nim#L109
func serializeRLNProof(proof *rln.RateLimitProof, epoch rln.Epoch, rlnIdentifier []byte) ([]byte, error) {
	rLimiProof := &rlnpb.RateLimitProof{
		Proof:         proof.Proof[:],
		MerkleRoot:    proof.MerkleRoot[:],
		Epoch:         epoch[:],
		ShareX:        proof.ShareX[:],
		ShareY:        proof.ShareY[:],
		Nullifier:     proof.Nullifier[:],
		RlnIdentifier: rlnIdentifier[:],
	}

	ser, err := proto.Marshal(rLimiProof)
	if err != nil {
		return nil, errors.Wrap(err, "error when marshalling proof")
	}

	return ser, nil
}
