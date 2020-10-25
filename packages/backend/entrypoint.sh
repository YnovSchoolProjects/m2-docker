#!/usr/bin/env bash

vars=("MONGODB_USERNAME" "MONGODB_PASSWORD" "MONGODB_ENDPOINT" "REDIS_DSN")

for var in vars ; do
    if [[ -n ${!var} ]]; then
      sed -i "s/${var}=(.*)/${var}=${!var}/g" .env
    fi
done

./app "$@"