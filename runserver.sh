#!/usr/bin/env bash
# Run the development server.

export API_GATE_ADDR=localhost:3001
go run ./app/gates/api/cmd/main.go
