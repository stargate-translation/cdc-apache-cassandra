#!/bin/bash

docker buildx build --platform linux/arm64 . \
  --push -t yarkhinephyo2019/dse-next:ayman-cdc \
  --build-arg COMMITMOG_SYNC_PERIOD_IN_MS=2000 \
  --build-arg CDC_TOTAL_SPACE_IN_MB=70
