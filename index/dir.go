package index

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"
)

type dirBasedIndex string

func IndexDirectory(path string) (Index, error) {
	if err := validateIndexDirectory(path); err != nil {
		return nil, err
	}

	return dirBasedIndex(path), nil
}

func validateIndexDirectory(path string) error {
	stats, err := os.Stat(path)
	if err != nil {
		return err
	}

	if !stats.IsDir() {
		return fmt.Errorf("%s is not a directory", path)
	}

	return nil
}

func (this dirBasedIndex) AddRecord(file File, record *Record) (*Record, error) {
	now := time.Now()
	record.LastWritten = now

	json, jsonErr := json.Marshal(&record)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if writeErr := ioutil.WriteFile(this.FullPath(file), json, 600); writeErr != nil {
		return nil, writeErr
	}

	return record, nil
}

func (this dirBasedIndex) FullPath(file File) string {
	return path.Join(string(this), string(file))
}

func (this dirBasedIndex) ContainsFile(file File) bool {
	_, err := os.Stat(this.FullPath(file))
	return err == nil
}