#!/usr/bin/env bash
# Execute the application database shell.
# Usage example: bin/dbshell.sh

# Load environment variables from .env file.
set -a; [ -f .env ] && . .env; set +a

# Execute a Mongo shell.
db/mongoshell.sh
