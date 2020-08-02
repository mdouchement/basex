package basex

import (
	"time"
)

// GenerateID creates a new ID based on time in nanoseconds.
func GenerateID() string {
	time.Sleep(3 * time.Nanosecond)
	return FormatBits(uint64(time.Now().UnixNano()), 62, false)
}
