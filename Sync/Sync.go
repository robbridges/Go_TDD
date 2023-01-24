package Sync

import "sync"

type Counter struct {
	Value int
	mu    sync.Mutex
}

// we need to use a mutex to lock and and unlock the count in the event someone tries to run this as go routines
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Value++

}

func (c *Counter) GetValue() int {
	return c.Value
}
