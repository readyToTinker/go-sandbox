package main

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	var name string
	namePrompt := &survey.Input{
		Message: "What's your name?",
	}
	err := survey.Ask([]*survey.Question{{
		Name:    "name",
		Prompt:  namePrompt,
		Validate: survey.Required,
	}}, &name)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Hello, %s!\n", name)
}