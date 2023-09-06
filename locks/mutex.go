package locks

import "sync"

// UseMutex locks the mutex then executes the function before unlocking the mutex again.
func UseMutex(mutex *sync.Mutex, function func()) {
	mutex.Lock()
	function()
	mutex.Unlock()
}

// UseRead locks the read mutex then executes the function before unlocking the read mutex again.
func UseRead(mutex *sync.RWMutex, function func()) {
	mutex.RLock()
	function()
	mutex.RUnlock()
}

// UseWrite locks the write mutex then executes the function before unlocking the write mutex again.
func UseWrite(mutex *sync.RWMutex, function func()) {
	mutex.Lock()
	function()
	mutex.Unlock()
}
