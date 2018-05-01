#!/usr/bin/env bash
# Run the development server.

export API_GATE_ADDR=localhost:3001 \
  UPSALE_SECRET_KEY=XgK8SELWvdY7HEG9JAbjKgJj39RJJMyq

go run ./app/gates/api/cmd/main.go
