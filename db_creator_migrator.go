package pgmigratotions

import (
	"fmt"
	"io/fs"
)

type DbConnectionSettings struct {
	User           string
	Password       string
	Host           string
	Port           string
	PrivateKeyFile *string
	CertFile       *string
	CaCertFile     *string
}

type MigrationSettings struct {
	MigrationsFS      fs.FS
	Schema            string
	MigrationsDirName string
}

func CreateDbAndMigrate(dbConSettings DbConnectionSettings, migrationSettings MigrationSettings, dbName string) error {
	return createDbAndMigrate(dbConSettings, migrationSettings, dbName, false)
}

func MustCreateDbAndMigrate(dbConSettings DbConnectionSettings, migrationSettings MigrationSettings, dbName string) error {
	return createDbAndMigrate(dbConSettings, migrationSettings, dbName, true)
}

func createDbAndMigrate(dbConSettings DbConnectionSettings, migrationSettings MigrationSettings, dbName string, mustCreate bool) error {
	instanceConStr := BuildConnectionString(dbConSettings.User, dbConSettings.Password, dbConSettings.Host, dbConSettings.Port, nil, dbConSettings.PrivateKeyFile, dbConSettings.CertFile, dbConSettings.CaCertFile)
	con, err := SqlDbConnection(instanceConStr)
	if err != nil {
		return fmt.Errorf("pgmigratotions: cannot open the connection to the database instance; %w;", err)
	}
	defer con.Close()
	if mustCreate {
		err = MustCreateDb(con, dbName)
	} else {
		err = CreateDb(con, dbName)
	}
	if err != nil {
		return fmt.Errorf("pgmigratotions: cannot create database %s; %w;", dbName, err)
	}
	dbConStr := BuildConnectionString(dbConSettings.User, dbConSettings.Password, dbConSettings.Host, dbConSettings.Port, &dbName, dbConSettings.PrivateKeyFile, dbConSettings.CertFile, dbConSettings.CaCertFile)
	dbCon, err := SqlDbConnection(dbConStr)
	if err != nil {
		return fmt.Errorf("pgmigratotions: cannot open the connection to the database %s; %w;", dbName, err)
	}
	err = MigrateUp(dbCon, migrationSettings.MigrationsFS, migrationSettings.Schema, migrationSettings.MigrationsDirName)
	if err != nil {
		return fmt.Errorf("pgmigratotions: cannot migrate database %s; %w;", dbName, err)
	}
	return nil
}
