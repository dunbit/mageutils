package targets

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

func (d *dirObject) Get(ref string) string {
	if ref == "RootDir" {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		return wd
	}
	if val, ok := d.Map[ref]; ok {
		return val
	}
	return ""
}

func (d *dirObject) GetDefault(ref string, defaultPath string) string {
	val := d.Get(ref)
	if val == "" {
		return defaultPath
	}
	return val
}

func (d *dirObject) Add(ref string, path string) {
	d.Map[ref] = path
}

var dirs = newDirObject()

// Dir ...
type Dir mg.Namespace

// Info ...
func (Dir) Info() {
	fmt.Println("### Dir Info ###")
	root := dirs.Get("RootDir")
	fmt.Println("Ref:", "RootDir", "\tPath:", root)
	for key, value := range dirs.Map {
		fmt.Println("Ref:", key, "\tPath:", value)
	}
}
