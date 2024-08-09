#!/bin/sh

DATABASE_URL="postgres://surl_login:@localhost:5432/surl?sslmode=disable"
migrate -database $DATABASE_URL -path db/migrations up
