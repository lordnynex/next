#!/usr/bin/env bash
# Run the development server.

# Load environment variables from .env file.
set -a; [ -f .env ] && . .env; set +a

go run app/gates/api/cmd/main.go
