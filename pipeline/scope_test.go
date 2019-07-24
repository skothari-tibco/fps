package pipeline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPath(t *testing.T) {

	path := getPath("math['sample'][0]")
	assert.Equal(t, path, ".math.sample[0]")
}

func TestPath2(t *testing.T) {

	path := getPath("math['sample']")
	assert.Equal(t, path, ".math.sample")
}
