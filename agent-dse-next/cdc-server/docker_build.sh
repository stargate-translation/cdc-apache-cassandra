#!/bin/bash

docker buildx build --platform linux/arm64,linux/amd64 . \
  --push -t yarkhinephyo2019/proxy:1.0
