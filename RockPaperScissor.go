package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("Welcome to my game!\n")
	var user_score int = 0
	var comp_score int = 0
	for { // For infinite loop
		var user_input int
		fmt.Println("Enter a choice (rock[0], paper[1], scissors[2]): ")
		fmt.Scanln(&user_input)
		rand.Seed(time.Now().UnixNano())
		selection := rand.Intn(3-0) + 0
		if user_input == selection {
			fmt.Printf("Both players selected %v. It's a tie!\n", user_input)
		} else if user_input == 0 {
			if selection == 2 {
				user_score += 1
				fmt.Println("Rock smashes scissors! You win!")
			} else {
				comp_score += 1
				fmt.Println("Paper covers rock! You lose.")
			}
		} else if user_input == 1 {
			if selection == 0 {
				user_score += 1
				fmt.Println("Paper covers rock! You win!")
			} else {
				comp_score += 1
				fmt.Println("Scissors cuts paper! You lose.")
			}
		} else if user_input == 2 {
			if selection == 1 {
				user_score += 1
				fmt.Println("Scissors cuts paper! You win!")
			} else {
				comp_score += 1
				fmt.Println("Rock smashes scissors! You lose.")
			}
		} else {
			fmt.Println(("Invalid input! try again"))
		}
		var play_again string
		fmt.Println("Play again? (y/n): ")
		fmt.Scanln(&play_again)
		if play_again != "y" {
			if user_score == comp_score {
				fmt.Printf("It's a tie! you both scored %v\n", user_score)

			} else if user_score > comp_score {
				fmt.Printf("Congratulations! You won\nyour score - %v\ncomputer's score - %v\n", user_score, comp_score)

			} else {
				fmt.Printf("Hard luck! You lose\nyour score - %v\ncomputer's score - %v\n", user_score, comp_score)
			}
			fmt.Println("Thanks for using our game :D, hope to see you soon")
			break
		}

	}
}
