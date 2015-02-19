package main

import (
	"os"
	"path/filepath"
)

const kay_meta_dir string = ".kay"
const kay_index_dir string = "index"

type KayDir string

func (kayDir KayDir) Content() string {
	return string(kayDir)
}

func (kayDir KayDir) Meta() string {
	return filepath.Join(kayDir.Content(), kay_meta_dir)
}

func (kayDir KayDir) Index() string {
	return filepath.Join(kayDir.Meta(), kay_index_dir)
}

func (kayDir KayDir) In() bool {
	_, err := os.Stat(kayDir.Meta())
	return err == nil
}

func (kayDir KayDir) Make() error {
	return os.MkdirAll(kayDir.Meta(), 0755)
}

func (kayDir KayDir) ContentFiles() ([]File, error) {
	return GetFilesFromDir(kayDir.Content())
}
