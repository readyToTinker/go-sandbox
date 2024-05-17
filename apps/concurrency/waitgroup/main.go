package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Starting multiple goroutines and waiting for them to finish with Waitgroups in the main routine.
func main()  {
	fmt.Println("Hello Concurrency!")
	fmt.Println("Number of CPUs:", runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(runtime.NumCPU())

	var counter = struct{
		count int
	}{
		count: 0,
	}
	
	for i := 0; i < runtime.NumCPU(); i++ {
		go func(wg *sync.WaitGroup, i int){
			fmt.Printf("Hello from a %v goroutine!\n",i)
			counter.count++
			wg.Done()
		}(&wg, i)
			
	}

	wg.Wait()
	
	fmt.Printf("Hello said at %v times", counter.count)
}