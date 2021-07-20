# Blockbook

## This repository is specificly for Coreblockchain

### try not to build other coins from this repository, because some main functionalities has been changed for the purpose of CoreBlockhain

*to build the debian package run this command:*

`make all-corecoin`

it gives you a `deb package` located in `build` folder.

The following methods are supported:



- [Status](/docs/api.md#status) ✅



-  [Get block hash](/docs/api.md#get-block-hash) ✅



- [Get transaction](/docs/api.md#get-transaction) ✅



-   [Get transaction specific](/docs/api.md#get-transaction-specific) ✅



-  [Get address](/docs/api.md#get-address) ✅



- [Get xpub](/docs/api.md#get-xpub) (not supported in ed448) 🚫



- [Get utxo](/docs/api.md#get-utxo) (only for bitcoin types) 🚫



- [Get block](/docs/api.md#get-block) ✅



-  [Send transaction](/docs/api.md#send-transaction) ✅



-   [Tickers list](/docs/api.md#tickers-list) (no fiat rate provider currently available for xcb) 🚫



-  [Tickers](/docs/api.md#tickers) (no fiat rate provider currently available for xcb) 🚫



-  [Balance history](/docs/api.md#balance-history) ✅

configurations and behaviours are kept untouched, so feel free to use the official docs for configurations and other things.

you can find the official readme file [here](README_ORG.md).



