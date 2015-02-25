package kaydir

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kklipsch/kay/wd"
)

type KayDir string

func Get(working wd.WorkingDirectory) (KayDir, error) {
	path := metaPath(working)
	if _, err := os.Stat(path); err != nil {
		return "", fmt.Errorf("This is not a kay directory.")
	}

	return KayDir(path), nil

}

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
