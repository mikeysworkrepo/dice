package main

import (
	"bufio"

	"fmt"
	"math/rand"

	"os"
	"strings"
)

func Dice() string {
	diceType := map[string]string{
		"1": "8",
		"2": "24",
	}

	fmt.Println("1. 8 sided dice")
	fmt.Println("2. 24 sided dice")

	reader := bufio.NewReader(os.Stdin)
	var choice string

	for {
		fmt.Println("Make your selection (1-2): ")
		choice, _ = reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		// check to see if input is valid
		if _, exists := diceType[choice]; exists {

			return choice

		} else {
			fmt.Println("Invalid input. Please enter a number 1-2")
		}

	}

}

func Roll(choice string) string {
	// fmt.Println(choice)
	if choice == "1" {
		// rand.Seed(time.Now().UnixNano())
		choices := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
		n := rand.Intn(9)

		return choices[n]
	}

	return choice
}

func PostRoll() string {
	rerollOptions := map[string]string{
		"1": "true",
		"2": "false",
	}
	fmt.Println("1. Reroll \n2. Change dice")
	reader := bufio.NewReader(os.Stdin)
	var rerollChoice string

	for {
		fmt.Println("Choose what to do next (1 or 2)")
		rerollChoice, _ = reader.ReadString('\n')
		rerollChoice = strings.TrimSpace(rerollChoice)

		if _, exists := rerollOptions[rerollChoice]; exists {
			return rerollChoice
		} else {
			fmt.Println("Invalid input. Select either option 1 or 2")
		}
	}

}
