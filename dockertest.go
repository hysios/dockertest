package dockertest

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
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
	dir, err := os.Getwd()
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
