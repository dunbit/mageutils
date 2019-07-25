package dep

import (
	"errors"
	"io"
	"os"
	"path"
	"strconv"

	"github.com/magefile/mage/sh"
	"gopkg.in/yaml.v2"
)

// Error definitions
var (
	ErrInvalidRepo = errors.New("Git invalid repo path")
	ErrInvalidOut  = errors.New("Git invalid out path")
)

var gitCmd = sh.RunCmd("git")

// Git ...
type Git struct {
	Repo   string `yaml:"repo"`
	Out    string `yaml:"out"`
	Branch string `yaml:"branch"`
	Depth  int    `yaml:"depth"`
	Single bool   `yaml:"single"`
	Hash   string `yaml:"hash"`
}

// Config ...
type Config struct {
	Git []Git `yaml:"git"`
}

// ReadConfig ...
func ReadConfig(reader io.Reader) (*Config, error) {
	config := &Config{}
	err := yaml.NewDecoder(reader).Decode(config)
	if err != nil {
		return nil, err
	}
	err = config.Validate()
	if err != nil {
		return nil, err
	}

	return config, nil
}

// Validate ...
func (c Config) Validate() error {
	for _, g := range c.Git {
		err := g.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

// Validate ...
func (g Git) Validate() error {
	if g.Repo == "" {
		return ErrInvalidRepo
	}

	if g.Out == "" {
		return ErrInvalidOut
	}

	return nil
}

// Ensure ...
func (g Git) Ensure() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	info, err := os.Stat(path.Join(wd, g.Out))
	if os.IsNotExist(err) {
		err := g.clone()
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	if info != nil && !info.IsDir() {
		return errors.New("expected path to be a dir, but got file")
	}

	err = g.checkout()
	if err != nil {
		return err
	}

	return nil
}

func (g Git) clone() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	args := []string{
		"clone",
	}

	if g.Depth > 0 {
		args = append(args, "--depth", strconv.Itoa(g.Depth))
	}

	if g.Single {
		args = append(args, "--single-branch")
	}

	if g.Branch != "" {
		args = append(args, "-b", g.Branch)
	}

	args = append(args, g.Repo, path.Join(wd, g.Out))

	return gitCmd(args...)
}

func (g Git) checkout() (rerr error) {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	defer func() {
		rerr = os.Chdir(wd)
	}()

	err = os.Chdir(path.Join(wd, g.Out))
	if err != nil {
		return err
	}

	err = gitCmd("diff", "--exit-code")
	if err != nil {
		return err
	}

	if g.Hash != "" {
		err = gitCmd("reset", "--hard", g.Hash)
		if err != nil {
			return err
		}
	}

	return nil
}
