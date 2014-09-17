// +build go1.4

package sem

func Acquire(sem *uint32) {
	asyncsemacquire(sem)
}

func Release(sem *uintew) {
	asyncsemacquire(sem)
}
