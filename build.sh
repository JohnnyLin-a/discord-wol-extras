#!/bin/bash
docker buildx build -t ghcr.io/johnnylin-a/discord-wol-extras --platform=linux/arm64,linux/arm/v7,linux/amd64 --push .