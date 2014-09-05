package x86

import (
	"testing"
)

func TestPause(t *testing.T) {
	for i := 0; i < 100; i++ {
		Pause()
	}
}
