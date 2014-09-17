// +build go1.3

package sem

import (
	_ "runtime"
)

func Acquire(sem *uint32) {
	semacquire(sem)
}

func Release(sem *uint32) {
	semrelease(sem)
}

//go:noescape
func semacquire(s *uint32)

//go:noescape
func semrelease(s *uint32)
