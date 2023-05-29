package main

import (
	"os/exec"
	"path"
)

// DeploymentVersion represents an artifact name and its last_passed version from
// the statistics table
type DeploymentVersion struct {
	Name              string
	LastPassedVersion string
}

type DeploymentMap = map[string]string

type NomadEngine struct {
	DirectoryPath string
	Address       string //includes IP:Port and http(s)
	// DeploymentMap map[string]bool
}

func (ne NomadEngine) Launch(name, version string) error {
	jobFileName := path.Join(ne.DirectoryPath, name+".hcl")
	// exec.Command("nomad", "job" ,"run" ,"-address" ,"http://10.250.101.4:4646",jobFileName)
	jobVersion := "version=" + version
	job := exec.Command("nomad", "job", "run", "-address", ne.Address, "-var", jobVersion, jobFileName)
	err := job.Run()

	return err
}

// func (ne NomadEngine) Destroy(name string) error {

// }

// func (ne NomadEngine) Revert(name, verison string) error {

// }

func (ne NomadEngine) generateDeploymentMap() {}

func (ne NomadEngine) reconfigure() {}

func (ne NomadEngine) configure() {}

// TODO uses stats table to do these

// func (ne NomadEngine) up(name string, version string) error {
// 	jobFileName := path.Join(ne.DirectoryPath, name+".hcl")
// 	// exec.Command("nomad", "job" ,"run" ,"-address" ,"http://10.250.101.4:4646",jobFileName)
// 	jobVersion := "version=" + version
// 	job := exec.Command("nomad", "job", "run", "-address", ne.Address, "-var", jobVersion, jobFileName)
// 	err := job.Run()

// 	return err
// }

// func (ne NomadEngine) down() error {

// }
