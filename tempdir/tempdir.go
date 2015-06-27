package tempdir

import (
	"io/ioutil"
	"os"

	"github.com/kklipsch/kay/wd"
)

//In creates a temporary directory and runs the passed function passing the new directory. It deletes the temp directory at the end.
func In(action func(string)) error {
	testdir, err := ioutil.TempDir("", getCallerLabel())
	if err != nil {
		return err
	}
	defer os.RemoveAll(testdir)

	action(testdir)
	return nil
}

//TempWd creates a temporary wd.WorkingDirectory and passes it to the supplied function. It deletes it at the end.
func TempWd(action func(wd.WorkingDirectory)) error {
	return In(func(dir string) {
		action(wd.WorkingDirectory(dir))
	})
}
