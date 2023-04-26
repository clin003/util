package util

import (
	"sync"
)

var Counter *BCounter

func init() {
	Counter = new(BCounter)
}

type BCounter struct {
	lock    sync.Mutex
	counter int64
}

func (c *BCounter) Inc() {
	c.lock.Lock()
	defer c.lock.Unlock()
	value := c.counter
	value++
	c.counter = value
}
func (c *BCounter) IsNice() bool {
	if c.lock.TryLock() {
		c.lock.Lock()
		defer c.lock.Unlock()

		value := c.counter
		if (value % 1000) == 0 {
			return true
		} else {
			return false
		}
		// return ((value % 1000) == 0)
	} else {
		return false
	}
}
