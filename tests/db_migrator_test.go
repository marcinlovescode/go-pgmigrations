package tests

import (
	"embed"
	"fmt"
	"strings"
	"testing"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"

	pgmigratotions "playground/go/pgmigrations"
)

//go:embed scripts/*.sql
var embedMigrations embed.FS

func TestMigrateUpRunsMigrations(m *testing.T) {
	//Arrange
	var id string
	expectedDbName := strings.ToLower(fmt.Sprintf("dbcreate_%d_m1", uuid.New().ID()))
	sqlStatement := `SELECT Id FROM sampletable where Id='1337';`
	instanceCon, _ := pgmigratotions.SqlDbConnection(BuildDbConnectionString(nil))
	defer instanceCon.Close()
	pgmigratotions.CreateDb(instanceCon, expectedDbName)
	dbInstanceCon, _ := pgmigratotions.SqlDbConnection(BuildDbConnectionString(&expectedDbName))
	defer dbInstanceCon.Close()
	//Act
	migrateErr := pgmigratotions.MigrateUp(dbInstanceCon, embedMigrations, "test", "scripts")
	//Assert
	row := dbInstanceCon.QueryRow(sqlStatement)
	err := row.Scan(&id)
	require.Nil(m, migrateErr)
	require.Nil(m, err)
	require.Equal(m, "1337", id)
}

func TestMigrateDownRemovesLastMigration(m *testing.T) {
	var id string
	expectedDbName := strings.ToLower(fmt.Sprintf("dbcreate_%d_m2", uuid.New().ID()))
	sqlStatement := `SELECT Id FROM sampletable where Id='1337';`
	instanceCon, _ := pgmigratotions.SqlDbConnection(BuildDbConnectionString(nil))
	defer instanceCon.Close()
	pgmigratotions.CreateDb(instanceCon, expectedDbName)
	dbInstanceCon, _ := pgmigratotions.SqlDbConnection(BuildDbConnectionString(&expectedDbName))
	defer dbInstanceCon.Close()
	pgmigratotions.MigrateUp(dbInstanceCon, embedMigrations, "test", "scripts")
	//Act
	migrateDownErr := pgmigratotions.MigrateDown(dbInstanceCon, embedMigrations, "test", "scripts")
	//Assert
	row := dbInstanceCon.QueryRow(sqlStatement)
	err := row.Scan(&id)
	require.Nil(m, migrateDownErr)
	require.Equal(m, "sql: no rows in result set", err.Error())
}

func TestMigrateDownRemovesAllMigrations(m *testing.T) {
	//Arrange
	var migrationsCount int
	expectedDbName := strings.ToLower(fmt.Sprintf("dbcreate_%d_m3", uuid.New().ID()))
	sqlStatement := `SELECT count(1) as migrationsCount FROM test_db_version`
	instanceCon, _ := pgmigratotions.SqlDbConnection(BuildDbConnectionString(nil))
	defer instanceCon.Close()
	pgmigratotions.CreateDb(instanceCon, expectedDbName)
	dbInstanceCon, _ := pgmigratotions.SqlDbConnection(BuildDbConnectionString(&expectedDbName))
	defer dbInstanceCon.Close()
	pgmigratotions.MigrateUp(dbInstanceCon, embedMigrations, "test", "scripts")
	//Act
	migrateDownErr := pgmigratotions.MigrateDownAll(dbInstanceCon, embedMigrations, "test", "scripts")
	//Assert
	row := dbInstanceCon.QueryRow(sqlStatement)
	err := row.Scan(&migrationsCount)
	require.Nil(m, migrateDownErr)
	require.Nil(m, err)
	require.Equal(m, 1, migrationsCount)
}
