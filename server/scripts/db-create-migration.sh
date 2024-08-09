#!/bin/sh

migrate create -ext sql -dir internal/database/migrations -seq create_users_table
