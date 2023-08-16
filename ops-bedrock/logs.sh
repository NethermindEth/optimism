#!/bin/bash

if [[ "$1" = "1" ]]; then
  docker logs -f "ops-bedrock-op-node-1" 2>&1 | grep --color -e 'Payload came from external builder.*id=0x\w\{64\}:\d\{1,\}' -e 'Skipping external builder this time.*reason=building with NoTxPool' -e '^'
fi

if [[ "$1" = "2" ]]; then
  docker-compose logs -f op-node-builder 2>&1
fi

if [[ "$1" = "3" ]]; then
  docker-compose logs -f l2 2>&1
fi

if [[ "$1" = "4" ]]; then
  docker logs -f "ops-bedrock-l2-builder-1" 2>&1 | grep --color -e 'Got best header.*id=0x\w\{64\}:\d\{1,\}' -e '^'
fi
