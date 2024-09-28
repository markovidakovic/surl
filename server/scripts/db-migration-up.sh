#!/bin/sh

DATABASE_URL="postgres://surl_login:Th9bcqiXcd9tYHlO@localhost:5432/surl?sslmode=disable"
migrate -database $DATABASE_URL -path internal/db/migrations up
