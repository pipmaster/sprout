Halo2 wasm template
==================

## Preparation
install `wasm-pack`
``` shell
curl https://rustwasm.github.io/wasm-pack/installer/init.sh -sSf | sh
```

install `solc`
Reference [install guid](https://docs.soliditylang.org/en/v0.8.9/installing-solidity.html)

## Build halo2 wasm program
1. get template 

``` shell
git clone git@github.com:machinefi/sprout.git && cd examples/halo2-circuits
```

2. build wasm

``` shell
wasm-pack build --target nodejs --out-dir pkg
```

you will find `xx_bg.wasm` in the `pkg` 

## Advanced
You can also develop your own halo2 circuit program.

1. Write a circuit according to the [halo2 development documentation](https://zcash.github.io/halo2/user/simple-example.html), and put the circuit file in `src/circuits`.
2. Replace the `TODO` in `src/lib.rs`.
3. Build wasm with `wasm-pack build --target nodejs --out-dir pkg`.

## Build executable file

``` shell
cargo build --release
```

## Generate verify smart contract

``` shell
target/release/halo2-circuit solidity
```
you will find `Verifier.sol` under the current folder. Or you can run `target/release/halo2-circuit solidity -f path/filename.sol`.

## Local verify proof
1. Get halo2 proof 
if you can send messages to znode successfully, then you can execute `ioctl ws message send --project-id 10001 --project-version "0.1" --data "{\"private_a\": 3, \"private_b\": 4}"` to obtain a halo2 proof, then put it in a file, like `halo2-proof.json`.

2. verify
`--proof` is proof file, and `--public` is the public input

``` shell
target/release/halo2-circuit verify --proof halo2-proof.json --public 900
```
