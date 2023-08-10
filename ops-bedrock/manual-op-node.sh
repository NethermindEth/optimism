#!/bin/sh
set -exu

exec ./op-node/bin/op-node \
    --l1=ws://127.0.0.1:8560 \
    --l2=http://127.0.0.1:8552 \
    --l2.jwt-secret=./ops-bedrock/test-jwt-secret.txt \
    --rollup.config=./.devnet/rollup.json \
    --rpc.addr=0.0.0.0 \
    --rpc.port=9543 \
    --rpc.enable-admin \
    --p2p.bootnodes=$1
