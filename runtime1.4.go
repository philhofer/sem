// +build go1.4

package sem

func Acquire(sem *uint32) {
	asyncsemacquire(sem)
}

func Release(sem *uintew) {
	asyncsemrelease(sem)
}

//go:noescape
func asyncsemrelease(s *uint32)

//go:noescape
func asyncsemacquire(s *uint32)
