#!/usr/bin/env bash
# Execute a Mongo shell or a script if a file provided.

docker-compose exec mongo mongo "$MONGO_INITDB_DATABASE" \
  --authenticationDatabase admin \
  -u "$MONGO_INITDB_ROOT_USERNAME" \
  -p "$MONGO_INITDB_ROOT_PASSWORD" \
  $1
