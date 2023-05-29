package main

import (
	"log"
	"os/exec"
	"path"

	"golang.org/x/sync/errgroup"
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

func (ne NomadEngine) LaunchTest(name, version string, dmap DeploymentMap) error {

	// var wg sync.WaitGroup
	var wg errgroup.Group

	// wg.Add(1)

	wg.Go(func() error {
		// defer wg.Done()
		return ne.lunch(name, version)

	})

	for n, v := range dmap {

		// wg.Add(1)

		// go func() {
		// 	defer wg.Done()
		// 	ne.lunch(n, v)
		// }()

		// wg.Go(func() error {
		// 	// defer wg.Done()
		// 	err := ne.lunch(n, v)
		// 	return err
		// })

		f := func() error {
			err := ne.lunch(n, v)
			return err
		}

		wg.Go(f)

	}

	// wg.Wait()
	if err := wg.Wait(); err != nil {
		log.Println(err)
		return err
	} else {
		return nil
	}
}

// func (ne NomadEngine) Destroy(name string) error {

// }

// func (ne NomadEngine) Revert(name, verison string) error {

// }

func (ne NomadEngine) lunch(name, version string) error {

	jobFileName := path.Join(ne.DirectoryPath, name+".hcl")
	jobVersion := "version=" + version
	job := exec.Command("nomad", "job", "run", "-address", ne.Address, "-var", jobVersion, jobFileName)
	err := job.Run()

	return err
}

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
