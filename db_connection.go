package pgmigratotions

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func SqlDbConnection(conStr string) (*sql.DB, error) {
	con, err := sql.Open("postgres", conStr)
	if err != nil {
		return nil, fmt.Errorf("pgmigratotions: can't open the connection; %w;", err)
	}
	return con, nil
}
