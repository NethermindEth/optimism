./op-node \
  --l1=ws://localhost:8546 \
  --l2=http://localhost:9501 \
  --l2.jwt-secret=../test-jwt-secret.txt \
  --rollup.config=../../.devnet/rollup.json \
  --rpc.addr=0.0.0.0 \
  --rpc.port=7500 \
  --rpc.enable-admin
