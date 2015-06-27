package kaydir

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kklipsch/kay/wd"
)

//KayDir is the .kay subfolder
type KayDir string

//Get will look in the current working directory for the .kay directory.  If it doesn't exist this is an error.
func Get(working wd.WorkingDirectory) (KayDir, error) {
	path := metaPath(working)
	if _, err := os.Stat(path); err != nil {
		return "", fmt.Errorf("This is not a kay directory.")
	}

	return KayDir(path), nil

}

//Make will create a .kay directory in the working directory.
func Make(in wd.WorkingDirectory) (KayDir, error) {
	path := metaPath(in)
	if err := os.MkdirAll(path, 0755); err != nil {
		return "", err
	}

	if err := os.MkdirAll(filepath.Join(path, "index"), 0755); err != nil {
		return "", err
	}

	return KayDir(path), nil
}

func metaPath(base wd.WorkingDirectory) string {
	return filepath.Join(string(base), ".kay")
}
