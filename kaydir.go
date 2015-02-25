package main

import (
	"os"
	"path/filepath"

	"github.com/kklipsch/kay/index"
)

const kayMetaDir string = ".kay"
const kayIndexDir string = "index"

type KayDir string

func (kayDir KayDir) Content() string {
	return string(kayDir)
}

func (kayDir KayDir) Meta() string {
	return filepath.Join(kayDir.Content(), kayMetaDir)
}

func (kayDir KayDir) Index() string {
	return filepath.Join(kayDir.Meta(), kayIndexDir)
}

func (kayDir KayDir) In() bool {
	_, err := os.Stat(kayDir.Meta())
	return err == nil
}

func (kayDir KayDir) Make() error {
	err := os.MkdirAll(kayDir.Meta(), 0755)
	if err != nil {
		return err
	}
	return os.MkdirAll(kayDir.Index(), 0755)
}

func (kayDir KayDir) ContentFiles() ([]index.File, error) {
	return GetFilesFromDir(kayDir.Content())
}
