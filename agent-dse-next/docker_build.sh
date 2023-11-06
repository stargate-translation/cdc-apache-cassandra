#!/bin/bash

docker buildx build --platform linux/arm64,linux/amd64 . \
  --push -t yarkhinephyo2019/dse-next:stargateio-1.0 \
  --build-arg COMMITMOG_SYNC_PERIOD_IN_MS=2000 \
  --build-arg CDC_TOTAL_SPACE_IN_MB=70
