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
}
func (c *BCounter) IsNice() bool {
	c.lock.Lock()
	defer c.lock.Unlock()

	value := c.counter
	if (value % 500) == 0 {
		return true
	} else {
		return false
	}
	// return c.IsNiceWithNum(500)
}
func (c *BCounter) IsNiceWithNum(num int64) bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	n := num
	if n <= 0 {
		n = 100
	}
	value := c.counter
	if (value % n) == 0 {
		return true
	} else {
		return false
	}
}
func (c *BCounter) Get() int64 {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.counter
}
