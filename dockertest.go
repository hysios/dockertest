package dockertest

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	Dockerfiles = []string{
		"docker-compose.yml",
		"docker-compose.yaml",
	}
)

func Prepare() error {
	dir, err := GoModuleRoot()
	if err != nil {
		return err
	}

	list, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, fi := range list {
		for _, s := range Dockerfiles {
			if fi.Name() == s && !fi.IsDir() {
				goto Found
			}
		}
	}

	return errors.New("not found dockerfiles")

Found:
	if IsDockerRunning() {
		return errors.New("always running")
	}

	return nil
}

func IsDockerRunning() bool {
	b, err := exec.Command("docker-compose", "ps", "-q", "|", "wc", "-l").Output()
	if err != nil {
		return false
	}

	n, err := strconv.Atoi(strings.TrimSpace(string(b)))
	if err != nil {
		return false
	}

	return n > 0
}

func RunDockerCompose() error {
	return exec.Command("docker-compose", "-d", "up").Run()
}

func GoModuleRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	paths := strings.Split(dir, string(filepath.Separator))

	l := len(paths)
	for i := 0; i < l; i++ {
		dir := strings.Join(paths[l-i:], string(filepath.Separator))

		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		} else if os.IsNotExist(err) {
			continue
		}
	}

	return "", os.ErrNotExist
}
