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

func pathToIndex(kd kaydir.KayDir) string {
	return filepath.Join(string(kd), "index")
}

//Make creates the index subfolder in the .kay directory
func Make(kd kaydir.KayDir) error {
	return os.MkdirAll(pathToIndex(kd), 0755)
}

//Get creates an in memory Index from the index subfolder of the .kay directory.  If index does not exist it is an error.
func Get(kd kaydir.KayDir) (Index, error) {
	path := pathToIndex(kd)

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

func (index dirBasedIndex) AddChapter(chap chapter.Chapter, record *Record) (*Record, error) {

	if index.ContainsChapter(chap) {
		return nil, fmt.Errorf("Index already contains %v", chap)
	}

	now := time.Now()
	record.LastWritten = now

	json, jsonErr := json.Marshal(&record)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if writeErr := ioutil.WriteFile(index.FullPath(chap), json, 0600); writeErr != nil {
		return nil, writeErr
	}

	return record, nil
}

func (index dirBasedIndex) FullPath(chap chapter.Chapter) string {
	return path.Join(string(index), string(chap))
}

func (index dirBasedIndex) ContainsChapter(chap chapter.Chapter) bool {
	_, err := os.Stat(index.FullPath(chap))
	return err == nil
}

func (index dirBasedIndex) AllIndexed() []chapter.Chapter {
	indexed := []chapter.Chapter{}

	files, err := ioutil.ReadDir(string(index))
	if err != nil {
		return indexed
	}

	for _, file := range files {
		indexed = append(indexed, chapter.Chapter(file.Name()))
	}

	return indexed
}

func (index dirBasedIndex) GetRecord(chap chapter.Chapter) (*Record, error) {
	file, readErr := ioutil.ReadFile(index.FullPath(chap))
	if readErr != nil {
		return nil, fmt.Errorf("Error reading record: %v", readErr)
	}

	var record Record
	if jsonErr := json.Unmarshal(file, &record); jsonErr != nil {
		return nil, fmt.Errorf("Error parsing record: %v (%s)", jsonErr, file)
	}

	return &record, nil
}
