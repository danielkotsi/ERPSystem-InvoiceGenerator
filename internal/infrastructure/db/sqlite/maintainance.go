package sqlite

import (
	"database/sql"
	"-invoice_manager/internal/models"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Database struct {
	DB     *sql.DB
	Config *models.Config
}

func NewDatabase(abspath string) *sql.DB {
	db := &Database{}
	db.Initialize(abspath)
	return db.DB
}

func (db *Database) Initialize(abspath string) {
	dbPath := filepath.Join(abspath, "newdata.db")

	isNewDB := false
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		log.Println("this is a new instance")
		file, err := os.Create(dbPath)
		if err != nil {
			log.Fatalf("failed to create database file: %v", err)
		}
		file.Close()
		isNewDB = true
		log.Println("Database file created successfully.")
	}

	newdb, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	newdb.Exec("PRAGMA foreign_keys=ON;")
	db.DB = newdb

	if isNewDB {
		migrationPath := filepath.Join(abspath, "assets", "migrations", "newmigration.sql")
		if err := db.runMigration(migrationPath); err != nil {
			log.Fatalf("migration failed: %v", err)
		}
	}

	// if err := db.CleanUpDatabase(); err != nil {
	// 	log.Println(err)
	// }

	log.Println("Database initialized successfully.")
}

func (db *Database) runMigration(path string) error {
	content, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return fmt.Errorf("failed to read migration file: %w", err)
	}

	// Διαχωρισμός των queries με ';'
	queries := strings.Split(string(content), ";")
	for _, query := range queries {
		query = strings.TrimSpace(query)
		if query == "" {
			continue
		}
		if _, err := db.DB.Exec(query); err != nil {
			return fmt.Errorf("failed to execute query: %w\nQuery: %s", err, query)
		}
	}
	log.Println("Migration executed successfully.")
	return nil
}
