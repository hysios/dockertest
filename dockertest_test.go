package dockertest

import (
	"testing"

	"github.com/tj/assert"
)

func TestPrepare(t *testing.T) {
	SetupDir("./tests/gosample/subdir")

	_, err := GoModuleRoot()
	assert.NoError(t, err)
}
