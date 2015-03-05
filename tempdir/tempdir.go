package tempdir

import (
	"io/ioutil"
	"os"

	"github.com/kklipsch/kay/wd"
)

func In(action func(string)) error {
	testdir, err := ioutil.TempDir("", getCallerLabel())
	if err != nil {
		return err
	} else {
		defer os.RemoveAll(testdir)
	}

	action(testdir)
	return nil
}

func TempWd(action func(wd.WorkingDirectory)) error {
	return In(func(dir string) {
		action(wd.WorkingDirectory(dir))
	})
}
