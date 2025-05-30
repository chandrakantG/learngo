package main

import "sync"

type Map struct {
	m map[int]int
	sync.Mutex
}

func (m *Map) Get(key int) (int, bool) {
	m.Lock()
	defer m.Unlock()
	i, ok := m.m[key]
	return i, ok
}
func (m *Map) Put(key, value int) {
	m.Lock()
	defer m.Unlock()
	m.m[key] = value
}
func (m *Map) Len() int {
	return len(m.m)
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	m := Map{m: make(map[int]int)}
	go func() {
		for i := 0; i < 100; i++ {
			m.Put(i, i)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 100; i++ {
			m.Len()
		}
		wg.Done()
	}()
	wg.Wait()
}

/**
A: Cannot compile
B: Can run without concurrency issues
C: Can run with concurrency issues => race condition occur at len func, so use RWMutex and use RLock and RUnlock in len func
D: panic
*/
