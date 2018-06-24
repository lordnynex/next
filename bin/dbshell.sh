#!/usr/bin/env bash
# Execute the application database shell.

# Move to the application root dir.
cd ..

# Load environment variables from .env file.
set -a; [ -f .env ] && . .env; set +a

# Execute a Mongo shell.
db/mongoshell.sh
