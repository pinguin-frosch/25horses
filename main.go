package main

import (
	"fmt"

	R "example.com/race"
)

func main() {
	fmt.Println("Racing horses!")

	race := R.New()

	var option string
	for {
		fmt.Println("\nT: Test winners")
		fmt.Println("R: Run a race")
		fmt.Println("Q: Quit")
		fmt.Println("S: Start again")

		fmt.Print("Choose an option: ")
		fmt.Scan(&option)

		if option == "R" {
			race.RunRace()
		} else if option == "T" {
			race.TestWinners()
		} else if option == "P" {
			race.PrintWinners()
		} else if option == "Q" {
			break
		}
	}
}
