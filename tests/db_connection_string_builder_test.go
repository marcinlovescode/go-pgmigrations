package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	pgmigratotions "playground/go/pgmigrations"
)

func TestBuildConnectionStringWithoutCertsReturnConnectionStringWithSSLModeDisable(t *testing.T) {
	//Arrange
	expectedUser := "postgresUser"
	expectedPassword := "postgresPassword"
	expectedHost := "localhost"
	expectedPort := "5432"
	expectedDatabase := "databasename"
	expectedSslMode := "disable"
	//Act
	conStr := pgmigratotions.BuildConnectionString(expectedUser, expectedPassword, expectedHost, expectedPort, &expectedDatabase, nil, nil, nil)
	//Assert
	require.Contains(t, conStr, fmt.Sprintf("user=%s", expectedUser))
	require.Contains(t, conStr, fmt.Sprintf("password=%s", expectedPassword))
	require.Contains(t, conStr, fmt.Sprintf("host=%s", expectedHost))
	require.Contains(t, conStr, fmt.Sprintf("port=%s", expectedPort))
	require.Contains(t, conStr, fmt.Sprintf("dbname='%s'", expectedDatabase))
	require.Contains(t, conStr, fmt.Sprintf("sslmode=%s", expectedSslMode))
}

func TestBuildConnectionStringWithCertsReturnConnectionStringWithSSLModeEnabled(t *testing.T) {
	//Arrange
	expectedUser := "postgresUser"
	expectedPassword := "postgresPassword"
	expectedHost := "localhost"
	expectedPort := "5432"
	expectedSslMode := "verify-ca"
	sslKey := "privateKey"
	sslcert := "publicCert"
	sslrootcert := "caCert"
	//Act
	conStr := pgmigratotions.BuildConnectionString(expectedUser, expectedPassword, expectedHost, expectedPort, nil, &sslKey, &sslcert, &sslrootcert)
	//Assert
	require.Contains(t, conStr, fmt.Sprintf("sslmode=%s", expectedSslMode))
	require.Contains(t, conStr, fmt.Sprintf("sslkey=%s", sslKey))
	require.Contains(t, conStr, fmt.Sprintf("sslcert=%s", sslcert))
	require.Contains(t, conStr, fmt.Sprintf("sslrootcert=%s", sslrootcert))
}

func TestBuildConnectionStringWithoutDbReturnConnectionStringWithoutDatabase(t *testing.T) {
	//Arrange
	expectedUser := "postgresUser"
	expectedPassword := "postgresPassword"
	expectedHost := "localhost"
	expectedPort := "5432"
	//Act
	conStr := pgmigratotions.BuildConnectionString(expectedUser, expectedPassword, expectedHost, expectedPort, nil, nil, nil, nil)
	//Assert
	require.NotContains(t, conStr, "dbname")
}
