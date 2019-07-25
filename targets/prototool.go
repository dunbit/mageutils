package targets

import (
	"fmt"
	"path"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var prototoolCmd = sh.RunCmd("prototool")

func init() {
	dirs.Add("ProtoDir", path.Join(dirs.Get("RootDir"), "proto"))
}

// Prototool ...
type Prototool mg.Namespace

// All ...
func (Prototool) All() {
	mg.Deps(Prototool.BreakCheck)
	mg.Deps(Prototool.Run)
	mg.Deps(Prototool.BreakUpdate)
}

// Info ...
func (Prototool) Info() error {
	fmt.Println("### Prototool Info ###")
	return prototoolCmd("version")
}

// Run ...
func (Prototool) Run() error {
	return prototoolCmd("all", dirs.GetDefault("ProtoDir", "./"))
}

// BreakCheck ...
func (Prototool) BreakCheck() error {
	return prototoolCmd("break", "check", "-f", path.Join(dirs.GetDefault("ProtoDir", "./"), "prototool.lock"), dirs.GetDefault("ProtoDir", "./"))
}

// BreakUpdate ...
func (Prototool) BreakUpdate() error {
	return prototoolCmd("break", "descriptor-set", "-o", path.Join(dirs.GetDefault("ProtoDir", "./"), "prototool.lock"), dirs.GetDefault("ProtoDir", "./"))
}
