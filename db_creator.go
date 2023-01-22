package pgmigratotions

import (
	"database/sql"
	"fmt"
	"strings"
)

func MustCreateDb(connection *sql.DB, dbName string) error {
	return createDb(connection, dbName, true)
}

func CreateDb(connection *sql.DB, dbName string) error {
	return createDb(connection, dbName, false)
}

func MustDropDb(connection *sql.DB, dbName string) error {
	return dropDb(connection, dbName, true)
}

func DropDb(connection *sql.DB, dbName string) error {
	return dropDb(connection, dbName, false)
}

func dbExists(connection *sql.DB, dbName string) (bool, error) {
	var dbExists int
	normalizedDbName := strings.ToLower(dbName)
	dbExistsSqlStatement := `SELECT COUNT(1) as dbExists FROM pg_database where datname =$1;`

	row := connection.QueryRow(dbExistsSqlStatement, normalizedDbName)
	err := row.Scan(&dbExists)
	if err != nil {
		return false, fmt.Errorf("pgmigratotions: cannot execute existence query; %w;", err)
	}
	return dbExists == 1, nil
}

func createDb(connection *sql.DB, dbName string, throwWhenExists bool) error {
	createDbSqlStatement := fmt.Sprintf("CREATE DATABASE %s", dbName)
	dbExists, err := dbExists(connection, dbName)
	if err != nil {
		return err
	}
	if dbExists && throwWhenExists {
		return fmt.Errorf("pgmigratotions: database: %s already exists; %w;", dbName, err)
	}
	if dbExists && !throwWhenExists {
		return nil
	}
	_, err = connection.Exec(createDbSqlStatement)
	if err != nil {
		return fmt.Errorf("pgmigratotions: cannot create database: %s; %w;", dbName, err)
	}
	return nil
}

func dropDb(connection *sql.DB, dbName string, throwWhenDoesNotExist bool) error {
	dropDbSqlStatement := fmt.Sprintf("DROP DATABASE %s", dbName)

	dbExists, err := dbExists(connection, dbName)
	if err != nil {
		return err
	}
	if !dbExists && throwWhenDoesNotExist {
		return fmt.Errorf("pgmigratotions: database: %s doesn't exist; %w;", dbName, err)
	}
	if !dbExists && !throwWhenDoesNotExist {
		return nil
	}
	_, err = connection.Exec(dropDbSqlStatement)
	if err != nil {
		return fmt.Errorf("pgmigratotions: cannot drop database: %s; %w;", dbName, err)
	}
	return nil
}
