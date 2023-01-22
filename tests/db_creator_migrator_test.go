package tests

import (
	"embed"
	"fmt"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	pgmigratotions "playground/go/pgmigrations"
)

//go:embed scripts/*.sql
var embedMigrationsScripts embed.FS

var defaultDbConnectionSettings pgmigratotions.DbConnectionSettings = pgmigratotions.DbConnectionSettings{
	User:           "postgres",
	Password:       "postgres",
	Host:           "localhost",
	Port:           "5432",
	PrivateKeyFile: nil,
	CertFile:       nil,
	CaCertFile:     nil,
}

var deafaultMigrationSettings pgmigratotions.MigrationSettings = pgmigratotions.MigrationSettings{
	MigrationsFS:      embedMigrationsScripts,
	Schema:            "test",
	MigrationsDirName: "scripts",
}

func TestMustCreateDbAndMigrateCreatesDatabaseAndRunsMigrations(m *testing.T) {
	//Arrange
	var id string
	expectedDbName := strings.ToLower(fmt.Sprintf("dbcreate_%d_cm1", uuid.New().ID()))
	sqlStatement := `SELECT Id FROM sampletable where Id='1337';`
	dbInstanceCon, _ := pgmigratotions.SqlDbConnection(BuildDbConnectionString(&expectedDbName))
	defer dbInstanceCon.Close()
	//Act
	migrateAndCreateErr := pgmigratotions.MustCreateDbAndMigrate(defaultDbConnectionSettings, deafaultMigrationSettings, expectedDbName)
	//Assert
	row := dbInstanceCon.QueryRow(sqlStatement)
	err := row.Scan(&id)
	require.Nil(m, migrateAndCreateErr)
	require.Nil(m, err)
	require.Equal(m, "1337", id)
}
