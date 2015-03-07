package index

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/kklipsch/kay/chapter"
	"github.com/kklipsch/kay/kaydir"
)

type dirBasedIndex string

func IndexPath(kd kaydir.KayDir) string {
	return filepath.Join(string(kd), "index")
}

func Make(kd kaydir.KayDir) error {
	return os.MkdirAll(IndexPath(kd), 0755)
}

func Get(kd kaydir.KayDir) (Index, error) {
	path := IndexPath(kd)

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

func (this dirBasedIndex) AddChapter(chap chapter.Chapter, record *Record) (*Record, error) {

	if this.ContainsChapter(chap) {
		return nil, fmt.Errorf("Index already contains %v", chap)
	}

	now := time.Now()
	record.LastWritten = now

	json, jsonErr := json.Marshal(&record)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if writeErr := ioutil.WriteFile(this.FullPath(chap), json, 0600); writeErr != nil {
		return nil, writeErr
	}

	return record, nil
}

func (this dirBasedIndex) FullPath(chap chapter.Chapter) string {
	return path.Join(string(this), string(chap))
}

func (this dirBasedIndex) ContainsChapter(chap chapter.Chapter) bool {
	_, err := os.Stat(this.FullPath(chap))
	return err == nil
}

func (this dirBasedIndex) GetRecord(chap chapter.Chapter) (*Record, error) {
	file, readErr := ioutil.ReadFile(this.FullPath(chap))
	if readErr != nil {
		return nil, readErr
	}

	var record Record
	if jsonErr := json.Unmarshal(file, &record); jsonErr != nil {
		return nil, jsonErr
	}

	return &record, nil
}