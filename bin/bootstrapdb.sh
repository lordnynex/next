#!/usr/bin/env bash
# Bootstrap the application database.
# Usage example: bin/bootstrapdb.sh

# Load environment variables from .env file.
set -a; [ -f .env ] && . .env; set +a

# Init the application database and create an admin.
db/mongoshell.sh /app/db/init.js
db/mongoshell.sh /app/db/createadmin.js
