#!/usr/bin/env bash
# Bootstrap the development Mongo database.

# Load environment variables from .env file.
set -a; [ -f .env ] && . .env; set +a

# Bootstrap the database and create an admin.
db/mongoscript.sh db/bootstrap.js
db/mongoscript.sh db/createadmin.js
