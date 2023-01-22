package pgmigratotions

import (
	"database/sql"
	"fmt"
	"github.com/pressly/goose/v3"
	"io/fs"
)

func MigrateUp(db *sql.DB, migrationsFS fs.FS, schema, migrationsDirName string) error {
	if err := setupMigrator(schema, migrationsFS); err != nil {
		return err
	}
	if err := goose.Up(db, migrationsDirName); err != nil {
		return fmt.Errorf("pgmigratotions: cannot migrate up; %w;", err)
	}
	return nil
}

func MigrateDown(db *sql.DB, migrationsFS fs.FS, schema, migrationsDirName string) error {
	if err := setupMigrator(schema, migrationsFS); err != nil {
		return err
	}
	if err := goose.Down(db, migrationsDirName); err != nil {
		return fmt.Errorf("pgmigratotions: cannot migrate down; %w;", err)
	}
	return nil
}

func MigrateDownAll(db *sql.DB, migrationsFS fs.FS, schema, migrationsDirName string) error {
	if err := setupMigrator(schema, migrationsFS); err != nil {
		return err
	}
	if err := goose.DownTo(db, migrationsDirName, 0); err != nil {
		return fmt.Errorf("pgmigratotions: cannot migrate down all; %w;", err)
	}
	return nil
}

func setupMigrator(schema string, migrationsFS fs.FS) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("pgmigratotions: cannot setup db dialect; %w;", err)
	}
	goose.SetTableName(fmt.Sprintf("%s_db_version", schema))
	goose.SetBaseFS(migrationsFS)
	return nil
}
