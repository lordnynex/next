#!/usr/bin/env bash
# Execute a Mongo script.

mongo "$UPSALE_MONGO_DATABASE" \
  --authenticationDatabase "$UPSALE_MONGO_DATABASE" \
  -u "$UPSALE_MONGO_USERNAME" \
  -p "$UPSALE_MONGO_PASSWORD" \
  "$1"
