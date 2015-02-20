package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func InTempDir(t *testing.T, label string, action func(string)) {
	testdir, err := ioutil.TempDir("", label)
	if err != nil {
		t.Fatalf("Could not create temp dir for %s: %v", label, err)
	} else {
		defer os.RemoveAll(testdir)
	}

	action(testdir)
}

func Assert(t *testing.T, shouldPass bool, msg string) {
	if !shouldPass {
		t.Errorf(msg)
	}
}

func FailIfError(t *testing.T, err error, msg string) {
	if err != nil {
		t.Fatalf("%s :%v", msg, err)
	}
}

func FailIfNoError(t *testing.T, err error, msg string) {
	if err == nil {
		t.Fatalf("No error:%s", msg)
	}
}

func AssertStringArray(t *testing.T, result []string, expected []string, msg string) {
	if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", expected) {
		t.Fatalf("Expected %v Got %v: %s", result, expected, msg)
	}
}

func AssertTags(t *testing.T, result []Tag, expected []Tag, msg string) {
	Assert(t, fmt.Sprintf("%v", result) == fmt.Sprintf("%v", expected), msg)
}

func AssertRecords(t *testing.T, record *record, expected *record, msg string) {
	Assert(t, record != nil, fmt.Sprintf("%v - sut is nil", record))
	Assert(t, expected != nil, fmt.Sprintf("%v - expected is nil", expected))
	Assert(t, record.Year == expected.Year, fmt.Sprintf("%s - Year: records don't match: %v, %v", msg, record, expected))
	Assert(t, record.File() == expected.File(), fmt.Sprintf("%s - File: records don't match: %v, %v", msg, record, expected))
	Assert(t, record.Note == expected.Note, fmt.Sprintf("%s - Note: records don't match: %v, %v", msg, record, expected))
	Assert(t, record.DateAdded == expected.DateAdded, fmt.Sprintf("%s - DateAdded: records don't match: %v, %v", msg, record, expected))
	AssertTags(t, record.Tags, expected.Tags, fmt.Sprintf("%s - Tags: records don't match: %v, %v", msg, record, expected))
}
