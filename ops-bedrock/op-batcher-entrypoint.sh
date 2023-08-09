#!/bin/sh
set -exu

curl \
    --fail \
    --retry 10 \
    --retry-delay 2 \
    --retry-connrefused \
    -X POST \
    -H "Content-Type: application/json" \
    --data '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["0x0", false],"id":1}' \
    "$OP_BATCHER_L1_ETH_RPC"

exec op-batcher
