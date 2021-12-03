package util

import (
	"math/rand"
	"time"
)

// 休眠随机时间（分钟）2_n*2
func SleepRandTimeMinute(n int64) {
	rand.Seed(time.Now().Unix())
	randInt := rand.Int63n(n)
	time.Sleep(time.Duration(randInt) * 2 * time.Minute)
}

// 休眠随机时间（秒） 2_n*2
func SleepRandTimeSecond(n int64) {
	rand.Seed(time.Now().Unix())
	randInt := rand.Int63n(n)
	time.Sleep(time.Duration(randInt) * 2 * time.Second)
}

// 休眠随机时间（分钟）2_n*2
func RandTimeMinute(n int64) time.Duration {
	rand.Seed(time.Now().Unix())
	randInt := rand.Int63n(n)
	return time.Duration(randInt) * 2 * time.Minute
}

// 休眠随机时间（秒） 2_n*2
func RandTimeSecond(n int64) time.Duration {
	rand.Seed(time.Now().Unix())
	randInt := rand.Int63n(n)
	return time.Duration(randInt) * 2 * time.Second
}
