package dep

import (
	"fmt"
	"os"
	"path"

	"github.com/dunbit/mageutils/pkg/dep"
	"github.com/dunbit/mageutils/targets/dir"
	"github.com/magefile/mage/mg"
)

// Deps ...
type Deps mg.Namespace

// Info ...
func (Deps) Info() error {
	file, err := os.Open(path.Join(dir.Get("RootDir"), "deps.yaml"))
	if err != nil {
		return err
	}

	config, err := dep.ReadConfig(file)
	if err != nil {
		return err
	}

	fmt.Println("### Deps Info ###")

	for _, git := range config.Git {
		fmt.Println("Repo:", git.Repo)
		fmt.Println(" Out:", git.Out)
		if git.Branch != "" {
			fmt.Println(" Branch:", git.Branch)
		}
		if git.Hash != "" {
			fmt.Println(" Hash:", git.Hash)
		}
	}

	return nil
}

// Example explains how to use deps.
func (Deps) Example() {
	fmt.Println("### Deps Example ###")
	fmt.Println("In order do use deps, create a config file in")
	fmt.Println("the root of your project: deps.yaml")
	fmt.Println("")
	fmt.Println("git:")
	fmt.Println("  - repo: http://path/to/repo.git")
	fmt.Println("    out: ./vendor/path/to/repo")
	fmt.Println("    branch: v1.2.0")
	fmt.Println("    single: true")
	fmt.Println("    depth: 1")
}
