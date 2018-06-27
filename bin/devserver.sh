#!/usr/bin/env bash
# Run the development server.
# Usage example: bin/devserver.sh

# Load environment variables from .env file.
set -a; [ -f .env ] && . .env; set +a

# Run the dev server.
go run app/gates/api/cmd/main.go
