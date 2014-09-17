package sem

import (
	"sync"
	"testing"
)

func TestSem(t *testing.T) {
	sem := new(uint32)
	*sem = 1

	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			Acquire(sem)
			Release(sem)
			wg.Done()
		}()
	}
	wg.Wait()
}
