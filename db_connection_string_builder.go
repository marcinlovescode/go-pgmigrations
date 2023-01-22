package pgmigratotions

import (
	"fmt"
	"strings"
)

func BuildConnectionString(user, password, host, port string, dbName, privateKeyFile, certFile, caCertFile *string) string {

	var connectionSB strings.Builder

	connectionSB.WriteString(fmt.Sprintf("user=%s password=%s host=%s port=%s",
		user,
		password,
		host,
		port))

	if dbName != nil {
		connectionSB.WriteString(fmt.Sprintf(" dbname='%s'", *dbName))
	}
	if privateKeyFile != nil && certFile != nil && caCertFile != nil {
		connectionSB.WriteString(fmt.Sprintf(" sslmode=verify-ca sslkey=%s sslcert=%s sslrootcert=%s",
			*privateKeyFile,
			*certFile,
			*caCertFile))
	} else {
		connectionSB.WriteString(" sslmode=disable")
	}
	return connectionSB.String()
}
