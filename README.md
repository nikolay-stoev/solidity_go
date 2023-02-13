https://www.ardanlabs.com/live-training-events/smart-contracts-with-go-feb-13-2023.html

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



