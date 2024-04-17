# go-waku-light

This repo contains a proof of concept for a `waku` light client, integrating [this](https://github.com/vacp2p/rln-contract/pull/31) modification of the RLN contract. It allows:
* Faster sync time. Uses `GetCommitments` instead of fetching events.
* Getting the Merkle root directly from the contract. No need to sync the whole tree.
* Getting the Merkle proof for any leaf directly from the contract. Not need to sync the whole tree.

For context, RLN is a decentralized rate-limiting protocol, that allows setting a limit on the number of messages sent by each entity, using zero-knowledge proofs. Said proofs prove that i) the member is whitelisted, without revealing its index, and ii) no more than 1 message is sent every epoch, which prevents double signaling.

The main motivation is to showcase how [this](https://github.com/vacp2p/rln-contract/pull/31) modification can help light clients become lighter. It makes proof generation and verification easier:
* Proof verification: Only requires the Merkle root, which is now available on-chain and can be fetched with a simple call. Before, one had to sync the whole Merkle tree with emitted events, and keep a local copy of it.
* Proof generation: Generating an RLN proof requires the Merkle proof of the leaf generating it. With this modification, said Merkle proof can be obtained directly from the contract.

Notes:
* Deployed in Polygon Layer 2 zkEVM, [see](https://cardona-zkevm.polygonscan.com/address/0x16abffcab50e8d1ff5c22b118be5c56f801dce54).
* This repo provides a simple CLI tool to showcase the functionalities.

## Usage

Build:
```
go build
```

Listen for new memberships and Merkle root changes. You can leave this running in another terminal.
```
./main listen
```

Register a new membership. You must provide a valid Polygon zkEVM account, see [faucet](https://faucet.polygon.technology/). Your membership information will be stored in a `.json` file.
```
./main register --priv-key=REPLACE_YOUR_PRIV_KEY
```

Fetches and logs the latest Merkle root and Merkle proof of the tree, using the contract as a source. Set `leaf-index` to your leaf index.
```
./main onchain-root
./main onchain-merkle-proof --leaf-index=1
```

Syncs the Merkle tree from the contract creating a local tree. The `chunk-size` indicates how many memberships are fetched at once. If too big, the RPC provider may error. Merkle proof can be locally computed, provide your `leaf-index`. Both `onchain` and `local` results should match.
```
./main local-root --chunk-size=500
./main local-merkle-proof --chunk-size=500 --leaf-index=1
```

Generates the RLN proof using a given membership, see the previous step. The proof is stored in a `.json` file. Message and epoch are hardcoded for simplicity. Note that the RLN can be generated `onchain` (which doesn't require to locally sync the tree since it uses the contract) and `local` (which syncs the tree locally).
```
./main onchain-generate-rln-proof --membership-file=membership_xxx.json
./main local-generate-rln-proof --membership-file=membership_xxx.json --chunk-size=500
```

Any RLN proof can be verified against the smart contract Merkle root.
```
./main verify-rln-proof --proof-file=proof_xxx.json
```


## Advanced

The `contract/contract.go` can be updated if the abi is changed as follows:

```
git clone https://github.com/ethereum/go-ethereum.git
cd go-ethereum
go build ./cmd/abigen
./abigen --abi=../go-waku-light/contract/abi.abi --pkg=contract --out=../go-waku-light/contract/contract.go
```