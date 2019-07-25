package dir

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
)

type dirObject struct {
	Map map[string]string
}

func newDirObject() *dirObject {
	return &dirObject{
		Map: make(map[string]string),
	}
}

// Get ...
func Get(ref string) string {
	if ref == "RootDir" {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		return wd
	}
	if val, ok := dirs.Map[ref]; ok {
		return val
	}
	return ""
}

// GetDefault ...
func GetDefault(ref string, defaultPath string) string {
	val := Get(ref)
	if val == "" {
		return defaultPath
	}
	return val
}

// Add ...
func Add(ref string, path string) {
	dirs.Map[ref] = path
}

var dirs = newDirObject()

// Dir ...
type Dir mg.Namespace

// Info ...
func (Dir) Info() {
	fmt.Println("### Dir Info ###")
	root := Get("RootDir")
	fmt.Println("Ref:", "RootDir", "\tPath:", root)
	for key, value := range dirs.Map {
		fmt.Println("Ref:", key, "\tPath:", value)
	}
}
