package util

import (
	"math"
	"sync"
)

// 计数器
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
	if value >= math.MaxInt64 {
		value = 0
	}
	c.counter = value
	// if c.lock.TryLock() {
	// 	c.lock.Lock()
	// 	defer c.lock.Unlock()
	// 	value := c.counter
	// 	value++
	// 	c.counter = value
	// } else {
	// 	fmt.Println(".", c.counter)
	// }
}
func (c *BCounter) IsNice() bool {
	c.lock.Lock()
	defer c.lock.Unlock()

	value := c.counter
	if (value % 1000) == 0 {
		return true
	} else {
		return false
	}

	// if c.lock.TryLock() {
	// 	c.lock.Lock()
	// 	defer c.lock.Unlock()

	// 	value := c.counter
	// 	if (value % 1000) == 0 {
	// 		return true
	// 	} else {
	// 		return false
	// 	}
	// 	// return ((value % 1000) == 0)
	// } else {
	// 	return false
	// }
}
func (c *BCounter) Get() int64 {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.counter
}
