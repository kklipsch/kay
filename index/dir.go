package index

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/kklipsch/kay/chapter"
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

func (this dirBasedIndex) AddChapter(chap chapter.Chapter, record *Record) (*Record, error) {
	now := time.Now()
	record.LastWritten = now

	json, jsonErr := json.Marshal(&record)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if writeErr := ioutil.WriteFile(this.FullPath(chap), json, 600); writeErr != nil {
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
