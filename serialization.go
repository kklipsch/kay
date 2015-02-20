package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func BuildIndex(indexDir string) (index, error) {
	if err := os.MkdirAll(indexDir, 0755); err != nil {
		return nil, err
	}

	files, geterr := GetFilesFromDir(indexDir)
	if geterr != nil {
		return nil, geterr
	}

	return readAllRecords(EmptyIndex(), indexDir, files)

}

func WriteIndex(indexDir string, index index) error {
	if err := os.MkdirAll(indexDir, 0755); err != nil {
		return err
	}

	// as we are going to be nuking the whole thing lets write it
	// first in case there are errors
	tempdir, temperr := createTempIndexDir(indexDir)
	if temperr != nil {
		return temperr
	}

	if err := writeAllRecords(index, tempdir); err != nil {
		return err
	}

	//need to nuke the whole index as a naive way to deal
	//with delete/move
	if removeErr := os.RemoveAll(indexDir); removeErr != nil {
		return removeErr
	}

	return os.Rename(tempdir, indexDir)
}

func DeserializeRecord(file File, index index, bytes []byte) (*record, error) {
	var rec record
	jsonErr := json.Unmarshal(bytes, &rec)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return index.SetRecord(&rec, file)
}

func readRecord(file File, path string, index index) (*record, error) {
	bytes, readErr := ioutil.ReadFile(filepath.Join(path, string(file)))
	if readErr != nil {
		return nil, readErr
	}

	return DeserializeRecord(file, index, bytes)
}

func readAllRecords(index index, path string, files []File) (index, error) {
	for _, file := range files {
		record, readErr := readRecord(file, path, index)
		if readErr != nil {
			return nil, readErr
		}
		index.SetRecord(record, file)
	}

	return index, nil
}

func createTempIndexDir(indexDir string) (string, error) {
	escapedName := strings.Replace(indexDir, "/", ".", -1)
	tempdir, temperr := ioutil.TempDir("", escapedName)
	if temperr != nil {
		return "", temperr
	} else {
		return tempdir, nil
	}
}

func SerializeRecord(record *record) ([]byte, error) {
	record.LastWritten = time.Now()
	return json.Marshal(record)
}

func writeRecord(record *record, tempdir string) error {
	jsonRecord, jsonerr := SerializeRecord(record)
	if jsonerr != nil {
		return jsonerr
	}

	indexFileName := path.Join(tempdir, string(record.File()))
	if writeerr := ioutil.WriteFile(indexFileName, jsonRecord, 0644); writeerr != nil {
		return writeerr
	}

	return nil
}

func writeAllRecords(index index, tempdir string) error {
	for _, record := range index.Records() {
		if err := writeRecord(record, tempdir); err != nil {
			return err
		}
	}

	return nil
}
