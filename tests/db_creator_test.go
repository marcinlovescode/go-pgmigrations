package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"

	pgmigratotions "playground/go/pgmigrations"
)

func TestCreateDbCreatesDatabase(m *testing.T) {
	//Arrange
	var dbExists string
	expectedDbName := strings.ToLower(fmt.Sprintf("dbcreate_%d_c1", uuid.New().ID()))
	sqlStatement := `SELECT COUNT(1) as dbExists FROM pg_database where datname =$1;`
	con, _ := pgmigratotions.SqlDbConnection(BuildDbConnectionString(nil))
	defer con.Close()
	//Act
	createDbErr := pgmigratotions.CreateDb(con, expectedDbName)
	//Assert
	row := con.QueryRow(sqlStatement, expectedDbName)
	err := row.Scan(&dbExists)
	require.Nil(m, createDbErr)
	require.Equal(m, "1", dbExists)
	require.Nil(m, err)
}

func TestCreateDbDoesNotThrowWhenDatabasebExists(m *testing.T) {
	//Arrange
	expectedDbName := strings.ToLower(fmt.Sprintf("dbcreate_%d_c2", uuid.New().ID()))
	con, _ := pgmigratotions.SqlDbConnection(BuildDbConnectionString(nil))
	defer con.Close()
	createDbErr := pgmigratotions.CreateDb(con, expectedDbName)
	//Act
	secondCallErr := pgmigratotions.CreateDb(con, expectedDbName)
	//Assert
	require.Nil(m, createDbErr)
	require.Nil(m, secondCallErr)
}

func TestMustCreateDbThrowsWhenDatabaseExists(m *testing.T) {
	//Arrange
	expectedDbName := strings.ToLower(fmt.Sprintf("dbcreate_%d_c3", uuid.New().ID()))
	con, _ := pgmigratotions.SqlDbConnection(BuildDbConnectionString(nil))
	defer con.Close()
	createDbErr := pgmigratotions.CreateDb(con, expectedDbName)
	//Act
	secondCallErr := pgmigratotions.MustCreateDb(con, expectedDbName)
	//Assert
	require.Nil(m, createDbErr)
	require.Contains(m, secondCallErr.Error(), fmt.Sprintf("pgmigratotions: database: %s already exists", expectedDbName))
}

func TestMustDropDbThrowsWhenDatabasebDoesNotExist(m *testing.T) {
	//Arrange
	expectedDbName := strings.ToLower(fmt.Sprintf("dbcreate_%d_c4", uuid.New().ID()))
	con, _ := pgmigratotions.SqlDbConnection(BuildDbConnectionString(nil))
	defer con.Close()
	//Act
	dropDbError := pgmigratotions.MustDropDb(con, expectedDbName)
	//Assert
	require.Contains(m, dropDbError.Error(), fmt.Sprintf("pgmigratotions: database: %s doesn't exist", expectedDbName))
}

func TestDropDbNotThrowsWhenDatabasebDoesNotExist(m *testing.T) {
	//Arrange
	expectedDbName := strings.ToLower(fmt.Sprintf("dbcreate_%d", uuid.New().ID()))
	con, _ := pgmigratotions.SqlDbConnection(BuildDbConnectionString(nil))
	defer con.Close()
	//Act
	dropDbError := pgmigratotions.DropDb(con, expectedDbName)
	//Assert
	require.Nil(m, dropDbError)
}

func TestDropDbDropsDatabase(m *testing.T) {
	//Arrange
	var dbExists string
	expectedDbName := strings.ToLower(fmt.Sprintf("dbcreate_%d_c5", uuid.New().ID()))
	sqlStatement := `SELECT COUNT(1) as dbExists FROM pg_database where datname =$1;`
	con, _ := pgmigratotions.SqlDbConnection(BuildDbConnectionString(nil))
	defer con.Close()
	createDbErr := pgmigratotions.CreateDb(con, expectedDbName)
	//Act
	dropDbErr := pgmigratotions.MustDropDb(con, expectedDbName)
	//Assert
	row := con.QueryRow(sqlStatement, expectedDbName)
	err := row.Scan(&dbExists)
	require.Nil(m, createDbErr)
	require.Nil(m, dropDbErr)
	require.Equal(m, "0", dbExists)
	require.Nil(m, err)
}
