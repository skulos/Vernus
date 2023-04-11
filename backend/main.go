// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	_ "github.com/mattn/go-sqlite3"
// )

// // Artifact represents the artifact information sent from Jenkins
// type Artifact struct {
// 	Name          string `json:"name"`
// 	Version       string `json:"version"`
// 	TestingStatus string `json:"testingStatus"`
// 	// Add more fields as needed
// }

// func main() {
// 	// Open a SQLite database
// 	db, err := sql.Open("sqlite3", "artifacts.db")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Create tables for each artifact name
// 	createTables(db)

// 	// Create a new Gin-Gonic router
// 	router := gin.Default()

// 	// Define a route for receiving artifact information from Jenkins
// 	router.POST("/deploy", func(c *gin.Context) {
// 		var artifact Artifact

// 		// Bind the JSON payload from Jenkins to the Artifact struct
// 		if err := c.ShouldBindJSON(&artifact); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// Validate the testingStatus field value
// 		allowedStatuses := map[string]bool{
// 			"Pending": true,
// 			"Passed":  true,
// 			"Failed":  true,
// 		}
// 		if _, ok := allowedStatuses[artifact.TestingStatus]; !ok {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid testingStatus value"})
// 			return
// 		}

// 		// Insert the artifact information into the corresponding table
// 		tableName := artifact.Name
// 		insertArtifact(db, tableName, artifact)

// 		// Send a response indicating successful receipt of the artifact information
// 		c.JSON(http.StatusOK, gin.H{"message": "Artifact received", "name": artifact.Name, "version": artifact.Version, "testingStatus": artifact.TestingStatus})
// 	})

// 	// Run the HTTP service
// 	if err := router.Run(":8080"); err != nil {
// 		log.Fatal(err)
// 	}
// }

// // createTables creates separate tables for each artifact name
// func createTables(db *sql.DB) {
// 	for _, name := range []string{"artifact1", "artifact2", "artifact3"} {
// 		query := fmt.Sprintf(`
// 			CREATE TABLE IF NOT EXISTS %s (
// 				id INTEGER PRIMARY KEY AUTOINCREMENT,
// 				name TEXT,
// 				version TEXT,
// 				testingStatus TEXT
// 			)
// 		`, name)

// 		if _, err := db.Exec(query); err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }

// // insertArtifact inserts the artifact information into the specified table
// func insertArtifact(db *sql.DB, tableName string, artifact Artifact) {
// 	query := fmt.Sprintf(`
// 		INSERT INTO %s (name, version, testingStatus) VALUES (?, ?, ?)
// 	`, tableName)

// 	if _, err := db.Exec(query, artifact.Name, artifact.Version, artifact.TestingStatus); err != nil {
// 		log.Fatal(err)
// 	}
// }

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// Artifact represents the artifact information sent from Jenkins
type Artifact struct {
	DateTime      time.Time
	Name          string `json:"name"`
	Version       string `json:"version"`
	TestingStatus string `json:"testingStatus"`
	// Add more fields as needed
}

type JenkinsArtifact struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func main() {

	// var jenkinsArtifact = struct {
	// 	Name    string `json:"name"`
	// 	Version string `json:"version"`
	// }{
	// 	Name:    "phoenix-testing",
	// 	Version: "v0.0.1",
	// }

	// json, _ := json.Marshal(jenkinsArtifact)

	// log.Println(string(json))

	// Open a SQLite database
	// db, err := sql.Open("sqlite3", "artifacts.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// artifact := Artifact{
	// 	Name:          "phoenix_testing",
	// 	Version:       "v0.0.2",
	// 	TestingStatus: "Pending",
	// 	DateTime:      time.Now(),
	// }

	// registerArtifact(db, artifact)

	// db.Close()

	db, err := sql.Open("sqlite3", "artifacts.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// tableName := "artifacts"
	// query := fmt.Sprintf(`
	// 		CREATE TABLE IF NOT EXISTS %s (
	// 			id INTEGER PRIMARY KEY AUTOINCREMENT,
	// 			dateTime DATETIME,
	// 			name TEXT,
	// 			version TEXT,
	// 			testingStatus TEXT
	// 		)
	// 	`, tableName)
	query := `
			CREATE TABLE IF NOT EXISTS artifacts (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				dateTime DATETIME,
				name TEXT,
				version TEXT,
				testingStatus TEXT
			)
		`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	// Create tables for each artifact name
	// createTables(db)

	// Create a new Gin-Gonic router
	router := gin.Default()

	// Define a route for receiving artifact information from Jenkins
	router.POST("/register", func(c *gin.Context) {
		var jenkinsArtifact JenkinsArtifact

		// Bind the JSON payload from Jenkins to the Artifact struct
		if err := c.ShouldBindJSON(&jenkinsArtifact); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			log.Println("error for jenkinsArtifact: ", err)
			return
		}

		// // Validate the testingStatus field value
		// allowedStatuses := map[string]bool{
		// 	"Pending": true,
		// 	"Passed":  true,
		// 	"Failed":  true,
		// }
		// if _, ok := allowedStatuses[artifact.TestingStatus]; !ok {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid testingStatus value"})
		// 	return
		// }

		// Register the artifact and insert version and testing status into the corresponding table
		artifact := Artifact{
			DateTime:      time.Now(),
			Name:          generateValidTableName(jenkinsArtifact.Name),
			Version:       jenkinsArtifact.Version,
			TestingStatus: "Pending",
		}

		err = registerArtifact(db, artifact)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Send a response indicating successful receipt of the artifact information
		c.JSON(http.StatusOK, gin.H{"message": "Artifact received", "time": artifact.DateTime, "name": artifact.Name, "version": artifact.Version, "testingStatus": artifact.TestingStatus})
	})

	// Run the HTTP service
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// registerArtifact registers a new artifact,
func registerArtifact(db *sql.DB, artifact Artifact) error {
	// Check if the artifact table exists
	tableName := artifact.Name
	query := fmt.Sprintf("SELECT name FROM sqlite_master WHERE type='table' AND name='%s'", tableName)
	var name string
	err := db.QueryRow(query).Scan(&name)
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
		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	// Insert the artifact version and testing status into the table
	query = fmt.Sprintf(`
		INSERT INTO %s (dateTime, version, testingStatus)
		VALUES (?, ?, ?)
	`, tableName)
	_, err = db.Exec(query, artifact.DateTime, artifact.Version, artifact.TestingStatus)
	if err != nil {
		log.Fatal(err)
		return err
	}

	artifactTableName := "artifacts"
	// Insert the artifact version and testing status into the table
	query = fmt.Sprintf(`
		INSERT INTO %s (dateTime, name, version, testingStatus)
		VALUES (?, ?, ?, ?)
	`, artifactTableName)
	_, err = db.Exec(query, artifact.DateTime, artifact.Name, artifact.Version, artifact.TestingStatus)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// generateValidTableName generates a valid table name based on the given name
func generateValidTableName(name string) string {
	// Remove any non-letter, non-number, and non-underscore characters
	// Replace spaces with underscores
	re := regexp.MustCompile(`[^a-zA-Z0-9_]+`)
	return re.ReplaceAllString(strings.Replace(name, " ", "_", -1), "")
}

// // createTables creates separate tables for each artifact name
// func createTable(db *sql.DB, artifact Artifact) {
// 	// for _, name := range []string{"artifact1", "artifact2", "artifact3"} {
// 	query := fmt.Sprintf(`
// 			CREATE TABLE IF NOT EXISTS %s (
// 				id INTEGER PRIMARY KEY AUTOINCREMENT,
// 				name TEXT,
// 				version TEXT,
// 				testingStatus TEXT
// 			)
// 		`, artifact.Name)

// 	if _, err := db.Exec(query); err != nil {
// 		log.Fatal(err)
// 	}
// 	// }
// }

// func getFirstPendingArtifactQuery(tableNames []string) string {
// 	query := "SELECT * FROM ("
// 	for i, tableName := range tableNames {
// 		if i > 0 {
// 			query += " UNION ALL "
// 		}
// 		query += fmt.Sprintf("SELECT name, version, testingStatus FROM %s", tableName)
// 	}
// 	query += fmt.Sprintf(") AS all_artifacts WHERE testingStatus = 'Pending' LIMIT 1;")
// 	return query
// }
