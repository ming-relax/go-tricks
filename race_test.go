package go_tricks

import (
	"testing"
	"sync"
)

func TestRace(t *testing.T) {
	var wg sync.WaitGroup

	foo := Foo{m: make(map[string]string)}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			foo.Set("a", "v")
		}()
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			foo.Get("a")
		}()
	}

	wg.Wait()
}
