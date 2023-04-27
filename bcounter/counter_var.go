package bcounter

var counter int64
var lock sync.Mutex

func Inc() {
	if lock.TryLock() {
		// lock.Lock()
		defer lock.Unlock()

		value := counter
		value++
		counter = value
	} else {
		fmt.Println(".", counter)
	}
}
func IsNice() bool {
	if lock.TryLock() {
		// lock.Lock()
		defer lock.Unlock()

		value := counter
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
