package main

import "fmt"

func main() {
	fmt.Println("Welcome to my shitty DND dice app! \nPlease select the dice you would like to roll..")

	fmt.Printf("you have rolled a %s!", Roll(Dice()))
	fmt.Scanln()

	if PostRoll() == "true" {
		fmt.Printf("you have rolled a %s!", Roll(Dice()))
		fmt.Scanln()
	} else {
		fmt.Println("placeholder for change dice")
	}

}
