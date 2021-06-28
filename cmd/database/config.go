package database

import (
	"fmt"
)

var (
	pguser  = "postgres"
	pgpass  = "postgres"
	pgport  = "5432"
	pghost  = "localhost"
	pgTable = "postgres"
	pgCS    = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		pghost, pgport, pguser, pgTable, pgpass)
)
