package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
)

type DatabaseHandler struct {
	db *sql.DB
}

func GenerateValidTableName(name string) string {
	return strings.Replace(name, "-", "_", -1)
}

func (dh DatabaseHandler) Connect() error {
	var err error
	dh.db, err = sql.Open("sqlite3", "artifacts.db")
	return err
}

func (dh DatabaseHandler) Close() error {
	err := dh.db.Close()
	return err
}

func (dh DatabaseHandler) CreateInitTables() {

	artifacts := fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	dateTime DATETIME,
	name TEXT,
	version TEXT,
	testingStatus TEXT
			)
		`, artefacts)

	_, err := dh.db.Exec(artifacts)
	if err != nil {
		log.Fatal(err)
	}

	services := fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	dateTime DATETIME,
	name TEXT,
	version TEXT,
	testingStatus TEXT
)
`, services)
	_, err = dh.db.Exec(services)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: stats table

}

// RegisterReleaseArtifact registers a new release artifact
func (dh DatabaseHandler) RegisterReleaseArtifact(artefact ReleaseArtifact) error {

	// Check if the artifact table exists
	tableName := artefact.Name
	query := fmt.Sprintf("SELECT name FROM sqlite_master WHERE type='table' AND name='%s'", tableName)
	var name string
	err := dh.db.QueryRow(query).Scan(&name)
	if err != nil && err != sql.ErrNoRows {
		log.Println(err)
		return err
	}

	// If the table does not exist, create it
	if err == sql.ErrNoRows {
		query := fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %s (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				dateTime DATETIME,
				version TEXT,
				testingStatus TEXT
			)
		`, tableName)
		_, err = dh.db.Exec(query)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	// Register with own table
	query = fmt.Sprintf(`
		INSERT INTO %s (dateTime, version, testingStatus)
		VALUES (?, ?, ?)
	`, tableName)
	_, err = dh.db.Exec(query, artefact.DateTime, artefact.Version, artefact.TestingStatus)

	if err != nil {
		log.Fatal(err)
		return err
	}

	// Insert the artifact version and testing status into the to test table
	query = fmt.Sprintf(`
		INSERT INTO %s (dateTime, name, version, testingStatus)
		VALUES (?, ?, ?, ?)
	`, artefacts)
	_, err = dh.db.Exec(query, artefact.DateTime, artefact.Name, artefact.Version, artefact.TestingStatus)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// func (dh DatabaseHandler) InsertReleaseArtifact() {}

func (dh DatabaseHandler) RemoveReleaseArtifact() (ReleaseArtifact, error) {
	var artifact ReleaseArtifact

	// Find the first artifact
	err := dh.db.QueryRow(`
		SELECT * FROM artifacts ORDER BY id LIMIT 1
	`).Scan(
		&artifact.ID,
		&artifact.DateTime,
		&artifact.Name,
		&artifact.Version,
		&artifact.TestingStatus,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return ReleaseArtifact{}, errors.New("no artifacts found")
		}
		return ReleaseArtifact{}, err
	}

	// Delete the artifact by ID
	query := `
		DELETE FROM artifacts WHERE id = ?
	`

	_, err = dh.db.Exec(query, artifact.ID)
	if err != nil {
		return ReleaseArtifact{}, err
	}

	return artifact, nil
}
