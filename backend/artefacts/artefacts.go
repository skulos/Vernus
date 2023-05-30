package artefacts

import (
	"strings"
	"time"
)

type VersionMatrix = map[string]string

// ReleaseArtifact represents the artifact information currently in the system
type ReleaseArtifact struct {
	ID            int
	DateTime      time.Time
	Name          string `json:"name"`
	Version       string `json:"version"`
	TestingStatus string `json:"testingStatus"`
	// Add more fields as needed
}

// NewRelease represents the information submitted by Jenkins to the system
type NewRelease struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func generateValidTableName(name string) string {
	return strings.Replace(name, "-", "_", -1)
}

func (nr NewRelease) ConvertToReleaseArtifact() ReleaseArtifact {
	return ReleaseArtifact{
		DateTime:      time.Now(),
		Name:          generateValidTableName(nr.Name),
		Version:       nr.Version,
		TestingStatus: "Pending",
	}
}
