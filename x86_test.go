package x86

import (
	"bytes"
	"crypto/rand"
	"io"
	"testing"
	"unsafe"
)

func TestPause(t *testing.T) {
	for i := 0; i < 100; i++ {
		Pause()
	}
}

func TestMemmoveNoAsm(t *testing.T) {
	const N = 1024

	first := make([]byte, N)
	second := make([]byte, N)

	io.ReadFull(rand.Reader, first)

	memmove(unsafe.Pointer(&second[0]), unsafe.Pointer(&first[0]), N)

	if !bytes.Equal(first, second) {
		t.Errorf("Bytes not equal: %x and %x", first, second)
	}
}

func TestMemmove(t *testing.T) {
	const N = 1024

	first := make([]byte, N)
	second := make([]byte, N)

	io.ReadFull(rand.Reader, first)

	Memmove(unsafe.Pointer(&second[0]), unsafe.Pointer(&first[0]), N)

	if !bytes.Equal(first, second) {
		t.Errorf("Bytes not equal: %x and %x", first, second)
	}
}
