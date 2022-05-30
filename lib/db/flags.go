package db

import "flag"

var (
	dbURL = flag.String("db-url", "postgres://user:pass@localhost:5432/cf", "database url")
)
