package dockertest

import (
	"testing"

	"github.com/tj/assert"
)

func TestGoModuleRoot(t *testing.T) {
	SetupDir("./tests/gosample/subdir")

	_, err := GoModuleRoot()
	assert.NoError(t, err)
}

func TestPrepare(t *testing.T) {
	SetupDir("./tests/gosample/subdir")

	err := Prepare()
	assert.NoError(t, err)
}
