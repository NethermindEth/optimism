#!/bin/bash

if [[ "$1" = "1" ]]; then
  docker logs -f "ops-bedrock-op-node-1" 2>&1 | grep --color -e 'Payload came from external builder.*id=0x\w\{64\}:\d\{1,\}' -e 'Payload from external builder failed' -e 'Skipping external builder this time' -e '^'
fi

if [[ "$1" = "2" ]]; then
  docker-compose logs -f mev-boost 2>&1
fi

if [[ "$1" = "3" ]]; then
  docker logs -f "ops-bedrock-l2-builder-1-1" 2>&1 | grep --color -e 'Got best header.*id=0x\w\{64\}:\d\{1,\}' -e '^'
fi

if [[ "$1" = "4" ]]; then
  docker logs -f "ops-bedrock-l2-builder-2-1" 2>&1 | grep --color -e 'Got best header.*id=0x\w\{64\}:\d\{1,\}' -e '^'
fi
