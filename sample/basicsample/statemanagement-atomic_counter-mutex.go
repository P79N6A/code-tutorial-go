/*
	Atomic Counters:
		The primary mechanism for managing state in Go is communication over channels.
		We saw this for example with worker pools.
		There are a few other options for managing state though.
		Here we’ll look at using the sync/atomic package for atomic counters accessed
		by multiple goroutines.
	Mutexes:
		In the previous example we saw how to manage simple counter state using atomic operations.
		For more complex state we can use a mutex to safely access data across multiple goroutines.
		
		
*/
package bsample

import "fmt"
import "time"
import "sync/atomic"
import "runtime"

import    "math/rand"
import    "sync"

/*
Running the program shows that we executed about 40,000 operations.
$ go run atomic-counters.go
ops: 40200

*/
func AtomicCounter() {
	// We’ll use an unsigned integer to represent our (always-positive) counter.
    var ops uint64 = 0
	// To simulate concurrent updates,
	// we’ll start 50 goroutines that each increment the counter about once a millisecond.
    for i := 0; i < 50; i++ {
        go func() {
            for {
				// To atomically increment the counter we use AddUint64,
				// giving it the memory address of our ops counter with the & syntax.
                atomic.AddUint64(&ops, 1)
				// Allow other goroutines to proceed.
                runtime.Gosched()
            }
        }()
    }
	// Wait a second to allow some ops to accumulate.
    time.Sleep(time.Second)
 	// In order to safely use the counter while it’s still being updated by other goroutines,
 	// we extract a copy of the current value into opsFinal via LoadUint64.
 	// As above we need to give this function the memory address &ops from which to fetch the value.
    opsFinal := atomic.LoadUint64(&ops)
    fmt.Println("ops:", opsFinal)
}

/*

 Running the program shows that we executed about 3,500,000 operations
 against our mutex-synchronized state.
$ go run mutexes.go
ops: 3598302
state: map[1:38 4:98 2:23 3:85 0:44]

*/
func MutexState() {
	// For our example the state will be a map.
    var state = make(map[int]int)
	// This mutex will synchronize access to state.
    var mutex = &sync.Mutex{}
	// To compare the mutex-based approach with another we’ll see later,
	// ops will count how many operations we perform against the state.
    var ops int64 = 0
	// Here we start 100 goroutines to execute repeated reads against the state.
    for r := 0; r < 100; r++ {
        go func() {
            total := 0
            for {
				// For each read we pick a key to access,
				// Lock() the mutex to ensure exclusive access to the state,
				// read the value at the chosen key, Unlock() the mutex,
				// and increment the ops count.
                key := rand.Intn(5)
                mutex.Lock()
                total += state[key]
                mutex.Unlock()
                atomic.AddInt64(&ops, 1)
				// In order to ensure that this goroutine doesn’t starve the scheduler,
				// we explicitly yield after each operation with runtime.Gosched().
				// This yielding is handled automatically
				// with e.g. every channel operation and
				// for blocking calls like time.Sleep,
				// but in this case we need to do it manually.
                runtime.Gosched()
            }
        }()
    }
	// We’ll also start 10 goroutines to simulate writes,
	// using the same pattern we did for reads.
    for w := 0; w < 10; w++ {
        go func() {
            for {
                key := rand.Intn(5)
                val := rand.Intn(100)
                mutex.Lock()
                state[key] = val
                mutex.Unlock()
                atomic.AddInt64(&ops, 1)
                runtime.Gosched()
            }
        }()
    }
	// Let the 10 goroutines work on the state and mutex for a second.
    time.Sleep(time.Second)
	// Take and report a final operations count.
    opsFinal := atomic.LoadInt64(&ops)
    fmt.Println("ops:", opsFinal)
	// With a final lock of state, show how it ended up.
    mutex.Lock()
    fmt.Println("state:", state)
    mutex.Unlock()
}