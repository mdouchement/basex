package basex_test

import (
	"testing"

	"github.com/mdouchement/basex"
	"github.com/stretchr/testify/assert"
)

func TestFormatBits(t *testing.T) {
	assert.Equal(t, "123456789", basex.FormatBits(123456789, 10, false))
	assert.Equal(t, "-18446744073586094827", basex.FormatBits(123456789, 10, true))

	assert.Equal(t, "726746425", basex.FormatBits(123456789, 8, false))
	assert.Equal(t, "75bcd15", basex.FormatBits(123456789, 16, false))
	assert.Equal(t, "8m0Kx", basex.FormatBits(123456789, 62, false))
}

func TestFormatBitsFromCharset(t *testing.T) {
	assert.Equal(t, "@#$%^&*()", basex.FormatBitsFromCharset("!@#$%^&*()", 123456789, 10, false))
	assert.Equal(t, "-@(%%&*%%!*$^(&!)%(#*", basex.FormatBitsFromCharset("!@#$%^&*()", 123456789, 10, true))

	assert.Equal(t, "*#&*%&%#^", basex.FormatBitsFromCharset("!@#$%^&*()", 123456789, 8, false))
}
