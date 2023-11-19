package disbursement

import (
	"sync"
)

var lockMutex sync.Mutex
var lockMap = make(map[int64]*sync.Mutex)

func getLock(key int64) (result *sync.Mutex) {
	lockMutex.Lock()
	defer lockMutex.Unlock()

	result, ok := lockMap[key]
	if !ok {
		result = &sync.Mutex{}
		lockMap[key] = result
	}

	return result
}

func removeLock(key int64) {
	lockMutex.Lock()
	defer lockMutex.Unlock()

	_, ok := lockMap[key]
	if ok {
		delete(lockMap, key)
	}
}
