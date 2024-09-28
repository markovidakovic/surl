#!/bin/sh

migrate create -ext sql -dir internal/db/migrations -seq create_users_table
