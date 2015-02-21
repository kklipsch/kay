package tempdir

import (
	"io/ioutil"
	"os"
)

func In(label string, action func(string)) error {
	testdir, err := ioutil.TempDir("", label)
	if err != nil {
		return err
	} else {
		defer os.RemoveAll(testdir)
	}
	action(testdir)
	return nil
}
