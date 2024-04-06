package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

func listenOnCommands(toR2, toR3 chan any) {
	// Open the keyboard for listening
	err := keyboard.Open()
	if err != nil {
		log.Fatal("Can't open keyboard, error: ", err)
	}
	defer keyboard.Close()

	// Listen for key events
	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch key {

		// close channels
		case keyboard.KeyCtrlC:
			handleClose(toR2, toR3)
			return

		// show help
		case keyboard.KeyCtrlH:
			handleHelp()

		// send command to R2
		case keyboard.KeyArrowLeft:
			handleSendMsgToR2(toR2)

		// send command to R3
		case keyboard.KeyArrowRight:
			handleSendMsgToR3(toR3)

		default:
			// Do nothing

		}
	}
}

func handleSendMsgToR3(toR3 chan any) {
	fmt.Println("Sending msg to R#3")
	toR3 <- "Hello R3"
}

func handleSendMsgToR2(toR2 chan any) {
	fmt.Println("Sending msg to R#2")
	toR2 <- "Hello R2"
}

func handleHelp() {
	help := `Help:
	- CTRL+H: this help 
	- CTRL+LeftArrow: send message to R#2 co-routine
	- CTRL+RightArrow: send message to R#3 co-routine
	- CTRL+C: close channels`
	fmt.Println(help)
}

// handle the closing of channels
func handleClose(toR1, toR2 chan any) {
	fmt.Println("Ctrl+C pressed, exiting...")
	close(toR1)
	close(toR2)
}
