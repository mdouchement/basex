package basex_test

import (
	"testing"

	"github.com/mdouchement/basex"
	"github.com/stretchr/testify/assert"
)

func TestGenerateID(t *testing.T) {
	assert.NotEqual(t, basex.GenerateID(), basex.GenerateID())
}
