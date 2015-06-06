package kayignore

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/kklipsch/kay/chapter"
	"github.com/kklipsch/kay/wd"
)

type ignorePattern string

type KayIgnore interface {
	Ignore(chap chapter.Chapter) bool
}

type ignoreNothing string

func (in ignoreNothing) Ignore(chap chapter.Chapter) bool {
	return false
}

const ignoreFile = ".kayignore"

func kayIgnorePath(workingDir wd.WorkingDirectory) string {
	return filepath.Join(string(workingDir), ignoreFile)
}

func Get(workingDir wd.WorkingDirectory) (KayIgnore, error) {
	filename := kayIgnorePath(workingDir)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return ignoreNothing(""), nil
	}

	fr, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	lines := []string{ignoreFile}
	scanner := bufio.NewScanner(fr)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		lines = append(lines, scanner.Text())
	}
	return patternList(lines), nil
}

func getPatterns(patterns ...string) KayIgnore {
	return patternList(patterns)
}

type patternList []string

func (pl patternList) Ignore(chap chapter.Chapter) bool {
	for _, pattern := range []string(pl) {
		if matched, _ := regexp.MatchString(pattern, string(chap)); matched {
			return true
		}
	}

	return false
}
