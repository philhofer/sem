// x86 provides
// access to some useful
// x86-64 stuff.
package x86

import (
	"unsafe"
)

//go:noescape
func Pause()

// no-op
func pause() {}

//go:noescape
func Memmove(dst unsafe.Pointer, src unsafe.Pointer, n uintptr)

func memmove(dst unsafe.Pointer, src unsafe.Pointer, n uintptr) {
	for i := uintptr(0); i < n; i++ {
		*(*byte)(unsafe.Pointer(uintptr(src) + i)) = *(*byte)(unsafe.Pointer(uintptr(dst) + i))
	}
}
