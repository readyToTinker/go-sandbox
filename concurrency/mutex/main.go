package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	counter int
}

const maxCount int = 40

var shared_value = Counter{maxCount}
var shared_value_with_mutex = Counter{maxCount}

// Using constructs from sync message to handle race conditions.
// sync.Mutex: handle race condition when updating values coming from shared resources, with Mutex we can lock the execution of a code piece for one routine only.
func main() {
	wg := sync.WaitGroup{}
	lock := sync.Mutex{}

	for i := 0; i < maxCount; i++ {
		wg.Add(1)
		go startRoutine(i, &wg)
		wg.Add(1)

		go startRoutineWithMutex(i, &wg, &lock)
	}
	wg.Wait()

	fmt.Println("Final number of increments without Mutex: " + fmt.Sprint(shared_value))
	fmt.Println("Final number of increments with Mutex: " + fmt.Sprint(shared_value_with_mutex))

}

func startRoutineWithMutex(i int, wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()
	lock.Lock()
	defer lock.Unlock()
	fmt.Printf("Go routine %v starting\n", i)
	fmt.Printf("Value before %v ", shared_value_with_mutex.counter)
	decreased := shared_value_with_mutex.counter - 1
	fmt.Printf("and value after %v\n", decreased)
	shared_value_with_mutex.counter = decreased

}

// All routine reads and writes the same shared variable with some delay, therefore the result will be random without Mutex.
func startRoutine(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Go routine %v starting\n", i)
	fmt.Printf("Value before %v ", shared_value.counter)
	decreased := shared_value.counter - 1
	fmt.Printf("and value after %v\n", decreased)
	shared_value.counter = decreased

}
