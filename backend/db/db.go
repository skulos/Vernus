package database

import (
	"Vernus/artefacts"
	"Vernus/nomad"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

const (
	artefact   string = "artefacts"
	services   string = "services"
	statistics string = "statistics"
)

type DatabaseHandler struct {
	Connection *sql.DB
}

// func (dh DatabaseHandler) Connect() error {
// 	var err error
// 	dh.Connection, err = sql.Open("sqlite3", "artifacts.db")
// 	return err
// }

func (dh DatabaseHandler) Close() error {
	err := dh.Connection.Close()
	return err
}

func (dh DatabaseHandler) CreateInitTables() {

	artifactsTable := fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	dateTime DATETIME,
	name TEXT,
	version TEXT,
	testingStatus TEXT
			)
		`, artefact)

	_, err := dh.Connection.Exec(artifactsTable)
	if err != nil {
		log.Fatal(err)
	}

	servicesTable := fmt.Sprintf(`
CREATE TABLE IF NOT EXISTS %s (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	dateTime DATETIME,
	name TEXT,
	version TEXT,
	testingStatus TEXT
)
`, services)
	_, err = dh.Connection.Exec(servicesTable)
	if err != nil {
		log.Fatal(err)
	}

	statisticsTable := fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		artifact_name TEXT UNIQUE,
		current_version TEXT,
		last_passed_version TEXT,
		last_failed_version TEXT
	)
`, statistics)

	_, err = dh.Connection.Exec(statisticsTable)
	if err != nil {
		log.Fatal(err)
	}

}

// RegisterReleaseArtifact registers a new release artifact
func (dh DatabaseHandler) RegisterReleaseArtifact(artifact artefacts.ReleaseArtifact) error {

	// Check if the artifact table exists
	tableName := artifact.Name
	query := fmt.Sprintf("SELECT name FROM sqlite_master WHERE type='table' AND name='%s'", tableName)
	var name string
	err := dh.Connection.QueryRow(query).Scan(&name)
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
		_, err = dh.Connection.Exec(query)
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
	_, err = dh.Connection.Exec(query, artifact.DateTime, artifact.Version, artifact.TestingStatus)

	if err != nil {
		log.Fatal(err)
		return err
	}

	// Insert the artifact version and testing status into the to test table
	query = fmt.Sprintf(`
		INSERT INTO %s (dateTime, name, version, testingStatus)
		VALUES (?, ?, ?, ?)
	`, artefact)
	_, err = dh.Connection.Exec(query, artifact.DateTime, artifact.Name, artifact.Version, artifact.TestingStatus)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// RemoveReleaseArtifact removes a release that has been tested form the waiting queue
func (dh DatabaseHandler) RemoveReleaseArtifact() (artefacts.ReleaseArtifact, error) {
	var artifact artefacts.ReleaseArtifact

	// Find the first artifact
	err := dh.Connection.QueryRow(`
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
			return artefacts.ReleaseArtifact{}, errors.New("no artifacts found")
		}
		return artefacts.ReleaseArtifact{}, err
	}

	// Delete the artifact by ID
	query := `
		DELETE FROM artifacts WHERE id = ?
	`

	_, err = dh.Connection.Exec(query, artifact.ID)
	if err != nil {
		return artefacts.ReleaseArtifact{}, err
	}

	return artifact, nil
}

// UpdateCurrentVersion updates the current version of a specific artifact in the statistics table
// If the artifact entry doesn't exist, it inserts a new entry with the provided version
func (dh DatabaseHandler) UpdateCurrentVersion(artifactName string, newVersion string) error {

	// Check if the artifact entry exists
	existenceQuery := fmt.Sprintf(`
		SELECT COUNT(*)
		FROM %s
		WHERE artifact_name = ?
	`, statistics)

	var count int
	err := dh.Connection.QueryRow(existenceQuery, artifactName).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {

		// Artifact entry doesn't exist, insert a new entry
		existenceQuery = `
			INSERT INTO statistics (artifact_name, current_version, last_passed_version, last_failed_version)
			VALUES (?, ?, ?, ?)
		`
		_, err = dh.Connection.Exec(existenceQuery, artifactName, newVersion, newVersion, newVersion)
		if err != nil {
			return err
		}

	} else {

		// Artifact entry exists, update the current version
		updateQuery := fmt.Sprintf(`
		UPDATE %s
		SET current_version = ?
		WHERE artifact_name = ?
	`, statistics)

		_, err := dh.Connection.Exec(updateQuery, newVersion, artifactName)
		if err != nil {
			return err
		}

	}

	return nil
}

// UpdateLastPassedVersion updates the last passed version version of a specific artifact in the statistics table
func (dh DatabaseHandler) UpdateLastPassedVersion(artifactName string, lastPassedVersion string) error {

	// Artifact entry exists, update the last passed version
	updateQuery := fmt.Sprintf(`
		UPDATE %s
		SET last_passed_version = ?
		WHERE artifact_name = ?
	`, statistics)

	_, err := dh.Connection.Exec(updateQuery, lastPassedVersion, artifactName)
	if err != nil {
		return err
	}

	return nil
}

// UpdateLastFailedVersion updates the last failed version of a specific artifact in the statistics table
func (dh DatabaseHandler) UpdateLastFailedVersion(artifactName string, lastFailedVersion string) error {

	// Artifact entry exists, update the last failed version
	updateQuery := fmt.Sprintf(`
		UPDATE %s
		SET last_failed_version = ?
		WHERE artifact_name = ?
	`, statistics)

	_, err := dh.Connection.Exec(updateQuery, lastFailedVersion, artifactName)
	if err != nil {
		return err
	}

	return nil
}

// GetDeploymentMap excluding a specific artifact and retrieves all of the artifact names and
// last_passed versions from the statistics table
func (dh DatabaseHandler) GetDeploymentMap(excludedArtifact string) (nomad.DeploymentMap, error) {

	query := fmt.Sprintf(`
	SELECT artifact_name, last_passed_version
	FROM %s
	WHERE artifact_name != ?
`, statistics)

	rows, err := dh.Connection.Query(query, excludedArtifact)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	dmap := make(nomad.DeploymentMap)

	for rows.Next() {
		var artifactName, lastPassedVersion string
		if err := rows.Scan(&artifactName, &lastPassedVersion); err != nil {
			return nil, err
		}

		dmap[artifactName] = lastPassedVersion

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return dmap, nil
}
