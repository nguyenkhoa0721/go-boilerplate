package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"go-boilerplate/cmd/cli/generator"
)

func main() {
	prompt := promptui.Select{
		Label: "Select",
		Items: []string{"Generate", "Exit"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if result == "Generate" {
		generator.Run()
	}
}
