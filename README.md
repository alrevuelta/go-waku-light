# go-waku-light

This repo contains a proof of concept of a `waku` light client, integrating [this](https://github.com/vacp2p/rln-contract/pull/31) modification in the RLN contract, which allows having the whole membership set + Merkle root on-chain. This makes light clients even lighter.

For context, RLN is a decentralized rate-limiting protocol, that allows setting a limit on the number of messages sent by each entity, using zero-knowledge proofs. Said proofs prove that i) the member is whitelisted, without revealing its index, and ii) no more than 1 message is sent every epoch, which prevents double signaling.

The main motivation is to showcase how [this](https://github.com/vacp2p/rln-contract/pull/31) modification can help light clients become lighter. It makes proof generation and verification easier:
* Proof verification: Only requires the Merkle root, which is now available on-chain and can be fetched with a simple call. Before, one had to sync the whole Merkle tree with emitted events, and keep a local copy of it.
* Proof generation: Requires the whole Merkle tree, which can be synced faster since this modification stores all the leaves in the contract. This approach is faster than syncing events, with the con of spending more gas on membership registration.

Notes:
* The new [contract](https://github.com/vacp2p/rln-contract/pull/31) uses a [BinaryIMT](https://github.com/privacy-scaling-explorations/zk-kit/blob/main/packages/imt.sol/contracts/BinaryIMT.sol).
* It is deployed in Polygon Layer 2 zkEVM, [see](https://testnet-zkevm.polygonscan.com/address/0x16aBFfCAB50E8D1ff5c22b118Be5c56F801Dce54).

This repo provides the following functionalities:
* Register an RLN membership in the contract in Layer 2 (Polygon zkEVM)
* Listen to new membership addition
* Get on-chain Merkle root
* Locally sync the membership Merkle tree (no need to sync events as before)
* Generate RLN proofs for a message
* Verify RLN proofs against the contract Merkle root.

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

Fetches and logs the latest Merkle root of the tree, using the contract as a source.
```
./main onchain-root
```

Syncs the Merkle tree from the contract. The `chunk-size` indicates how many memberships are fetched at once. If too big, the RPC provider may error.
```
./main sync-tree --chunk-size=500
```

Generates a proof using a given membership. The proof is stored in a `.json` file.
```
./main generate-proof --membership-file=membership_xxx.json
```

Verifies a given proof against the smart contract Merkle root.
```
./main verify-proof --proof-file=proof_xxx.json
```