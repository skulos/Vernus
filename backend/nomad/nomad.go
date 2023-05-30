package nomad

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

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

func (ne NomadEngine) DestroyTest() error {

	fileNames, err := ne.readFileNamesInDirectory()

	if err != nil {
		return err
	}

	for _, n := range fileNames {
		err := ne.destroy(n)

		if err != nil {
			return err
		}
	}

	return nil

	// var wg errgroup.Group

	// for i := 0; i < len(fileNames); i++ {

	// 	fmt.Println(i, ":", fileNames[i])

	// 	f := func() error {
	// 		err := ne.destroy(fileNames[i])
	// 		return err
	// 	}

	// 	wg.Go(f)

	// }

	// waitingError := wg.Wait()

	// if waitingError != nil {
	// 	log.Println(waitingError)
	// 	return waitingError
	// } else {
	// 	return nil
	// }
}

// func (ne NomadEngine) Revert(name, verison string) error {

// }

func (ne NomadEngine) lunch(name, version string) error {

	jobFileName := path.Join(ne.DirectoryPath, name+".hcl")
	jobVersion := "version=" + version
	job := exec.Command("nomad", "job", "run", "-address", ne.Address, "-var", jobVersion, jobFileName)
	err := job.Run()

	return err
}

func (ne NomadEngine) destroy(name string) error {

	job := exec.Command("nomad", "job", "stop", "-purge", "-address", ne.Address, "-namespace", "echo", name)
	err := job.Run()

	fmt.Println("Destroy err, ", name)

	return err
}

// ReadFileNamesInDirectory reads only the names of files in a directory and strips ".hcl" from the names
func (ne NomadEngine) readFileNamesInDirectory() ([]string, error) {
	fileNames := []string{}

	// Read the directory entries
	entries, err := os.ReadDir(ne.DirectoryPath)
	if err != nil {
		return fileNames, fmt.Errorf("failed to read directory: %v", err)
	}

	// Iterate over the directory entries
	for _, entry := range entries {
		if entry.IsDir() {
			// Skip directories
			continue
		}

		// Get the file name without the extension
		fileName := strings.TrimSuffix(entry.Name(), ".hcl")

		// Append the file name to the fileNames slice
		fileNames = append(fileNames, fileName)
	}

	return fileNames, nil
}

// func (ne NomadEngine) reconfigure() {}

// func (ne NomadEngine) configure() {}

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
