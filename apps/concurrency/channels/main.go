package main

import (
	"fmt"
	"sync"
)

func log_msg(id int, t ...any) {
	prefix := "R#" + fmt.Sprint(id) + ": "
	fmt.Print(append([]any{prefix}, t...)...)
	fmt.Print("\n")

}

func setupRoutine(id int, fromMain, between chan any, wg *sync.WaitGroup) {

	// start listener
	wg.Add(1)
	log_msg(id, "Starting")
	defer wg.Done()
	var idOther string
	if id == 2 {
		idOther = "R#3"
	} else {
		idOther = "R#2"
	}

	for {
		select {
		case cmd, ok := <-fromMain:
			if !ok {
				log_msg(id, "Channel closed: ", cmd)
				return
			}
			log_msg(id, "Command arrived from main: ", cmd)
			between <- cmd

		case msg := <-between:
			log_msg(id, "Message arrived from ", idOther, ": ", msg)

		}
	}
}

func main() {
	wg := sync.WaitGroup{}
	main_to_r2 := make(chan any)
	main_to_r3 := make(chan any)
	r2_to_r3 := make(chan any)

	fmt.Println("Show help with Ctrl+H")

	// starting R2 and R3 co-routines which are communicating with each other via a channel
	go setupRoutine(2, main_to_r2, r2_to_r3, &wg)
	go setupRoutine(3, main_to_r3, r2_to_r3, &wg)

	// keyboard monitoring
	listenOnCommands(main_to_r2, main_to_r3)
	wg.Wait()
}
