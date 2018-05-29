#!/usr/bin/env bash
# Run the development server.

export \
  UPSALE_MAIL_HOST=smtp.yandex.ru \
  UPSALE_MAIL_PORT=587 \
  UPSALE_MAIL_USERNAME=sail.notification@yandex.ru \
  UPSALE_MAIL_PASSWORD=q1w2e3asd \
  \
  UPSALE_MONGO_ADDRS=localhost:27017 \
  UPSALE_MONGO_DATABASE=demodev \
  UPSALE_MONGO_USERNAME=demodev \
  UPSALE_MONGO_PASSWORD=123

go run ./app/gates/api/cmd/main.go
