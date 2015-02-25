package kaydir

import (
	"fmt"
	"os"
	"path/filepath"
)

type KayDir string

func Get(start string) (KayDir, error) {
	path := metaPath(start)
	if _, err := os.Stat(path); err != nil {
		return "", fmt.Errorf("This is not a kay directory.")
	}

	return KayDir(path), nil

}

func Make(in string) (KayDir, error) {
	path := metaPath(in)
	if err := os.MkdirAll(path, 0755); err != nil {
		return "", err
	}

	if err := os.MkdirAll(filepath.Join(path, "index"), 0755); err != nil {
		return "", err
	}

	return KayDir(path), nil
}

func metaPath(base string) string {
	return filepath.Join(base, ".kay")
}
