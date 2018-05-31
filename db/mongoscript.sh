#!/usr/bin/env bash
# Execute a Mongo script.

mongo "$NEXT_MONGO_DATABASE" \
  --authenticationDatabase "$NEXT_MONGO_DATABASE" \
  -u "$NEXT_MONGO_USERNAME" \
  -p "$NEXT_MONGO_PASSWORD" \
  "$1"
