Notes for: https://www.ardanlabs.com/live-training-events/smart-contracts-with-go-feb-13-2023.html

Solidity smart contracts with Go tooling
Runs Geth in dev mode.

Use makefile for dependencies.
`make dev.setup`
Spin up a local env
`make geth-up`

All config goes into `/zarf`

`make geth-attach`
To open up JS console environment for making geth related API calls. For example
`eth.getBalance("0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd")` should give you the balance of the Coinbase account.

`make geth-deposit` 
to send eth to a few other accounts defined in `zarf/keystore`

-------------

Contracts:

`Basic`

```
public Items map (string -> uint256)
```

Can't iterate over maps - it would be costly. 
Use events - get it from tx 
receipt. 

```
event ItemSet(string key, uint256 value);
```

*Hacky - can we store immutable data in the event itself so we  save gas?

Generate `abi` and `bin`:
`make basic-build`
Creates a Golang package that we can use to deploy / interact with the smart contract.

`make basic-deploy`

```
% make basic-deploy
CGO_ENABLED=0 go run app/basic/cmd/deploy/main.go

Input Values
----------------------------------------------------
fromAddress: 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd
oneETHToUSD: 1487.5944560205614
oneUSDToETH: 0.0006722262213016597

Transaction Details
----------------------------------------------------
hash            : 0xbd6dfb42bfe29c9f5ed60c1665172b78b8e810409f316d0c3f9feb4d33dbcf10
nonce           : 1
gas limit       : 1600000
gas offer price : 0.875000001 GWei
value           : 0 GWei
max gas price   : 1400000.002 GWei
max gas price   : 2.08 USD

Contract Details
----------------------------------------------------
contract id     : 0x6dB8BD2Fd59009cF8aBB62D0883B5c200333666A

Waiting Logs
----------------------------------------------------
t=2023-02-13T21:46:49+0200 lvl=trce msg="Handled RPC response" reqid=6 duration="1.974µs"
t=2023-02-13T21:46:49+0200 lvl=trce msg="Transaction not yet mined" hash=0xbd6dfb42bfe29c9f5ed60c1665172b78b8e810409f316d0c3f9feb4d33dbcf10
t=2023-02-13T21:46:50+0200 lvl=trce msg="Handled RPC response"      reqid=7 duration="1.887µs"

Receipt Details
----------------------------------------------------
status          : 1
gas used        : 413926
gas price       : 0.875000001 GWei
gas price       : 0.00 USD
final gas cost  : 362185.2504 GWei
final gas cost  : 0.54 USD

Logs
----------------------------------------------------
t=2023-02-13T21:46:50+0200 lvl=trce msg="Handled RPC response"      reqid=8 duration="2.162µs"

Balance
----------------------------------------------------
balance before  : 1.157920892e+68 GWei
balance after   : 1.157920892e+68 GWei
balance diff    : 317077.278 GWei
balance diff    : 0.47 USD
```

`make basic-write`

CGO_ENABLED=0 go run app/basic/cmd/write/main.go

Input Values
----------------------------------------------------
fromAddress: 0x6327A38415C53FFb36c11db55Ea74cc9cB4976Fd
oneETHToUSD: 1490.0230939405994
oneUSDToETH: 0.0006711305375511619
contractID: 0x6dB8BD2Fd59009cF8aBB62D0883B5c200333666A
version: 1.1

Transaction Details
----------------------------------------------------
hash            : 0x71d5f1c1e8689dcf88596f8b5af0867d8a2a2b4097bff5903d1b2f3942f95f0c
nonce           : 2
gas limit       : 250000
gas offer price : 0.766024068 GWei
value           : 0 GWei
max gas price   : 191506.017 GWei
max gas price   : 0.29 USD

Receipt Details
----------------------------------------------------
status          : 1
gas used        : 47789
gas price       : 0.766024068 GWei
gas price       : 0.00 USD
final gas cost  : 36607.52419 GWei
final gas cost  : 0.05 USD

Logs
----------------------------------------------------
ItemSet
map[key:bill value:1000000]

Balance
----------------------------------------------------
balance before  : 1.157920892e+68 GWei
balance after   : 1.157920892e+68 GWei
balance diff    : 32360.35017 GWei
balance diff    : 0.05 USD