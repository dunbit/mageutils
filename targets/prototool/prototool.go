package prototool

import (
	"fmt"
	"path"

	"github.com/dunbit/mageutils/targets/dir"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var prototoolCmd = sh.RunCmd("prototool")

func init() {
	dir.Add("ProtoDir", path.Join(dir.Get("RootDir"), "proto"))
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
	return prototoolCmd("all", dir.GetDefault("ProtoDir", "./"))
}

// BreakCheck ...
func (Prototool) BreakCheck() error {
	return prototoolCmd("break", "check", "-f", path.Join(dir.GetDefault("ProtoDir", "./"), "prototool.lock"), dir.GetDefault("ProtoDir", "./"))
}

// BreakUpdate ...
func (Prototool) BreakUpdate() error {
	return prototoolCmd("break", "descriptor-set", "-o", path.Join(dir.GetDefault("ProtoDir", "./"), "prototool.lock"), dir.GetDefault("ProtoDir", "./"))
}
