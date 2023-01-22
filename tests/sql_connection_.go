package tests

import (
	_ "github.com/lib/pq"

	pgmigratotions "playground/go/pgmigrations"
)

func BuildDbConnectionString(dbName *string) string {
	return pgmigratotions.BuildConnectionString("postgres", "postgres", "localhost", "5432", dbName, nil, nil, nil)
}
