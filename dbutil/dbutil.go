package dbutil

import "github.com/lambda-platform/lambda/config"

// IsOracle returns true if the current database connection is Oracle.
func IsOracle() bool {
	return config.Config.Database.Connection == "oracle"
}

// IsPostgres returns true if the current database connection is PostgreSQL.
func IsPostgres() bool {
	return config.Config.Database.Connection == "postgres"
}
