package go_tricks

import "sync"

type Foo struct {
	sync.Mutex
	m map[string]string
}


func (foo Foo) Get(name string) string {
	foo.Lock()
	defer foo.Unlock()

	return foo.m[name]
}

func (foo *Foo) Set(name string, value string) {
	foo.Lock()
	defer foo.Unlock()

	foo.m[name] = value
}