package main

import (
	"fmt"
	"strings"
)

func main() {
	answer := "hello"
	attempts := 0
	correct_chars := 0

	var text_from_input string
	var guess string

	for attempts < 3 {
		fmt.Print("Guess a 5 letter word: ")
		fmt.Scan(&text_from_input)

		guess = strings.ToLower(text_from_input)
		fmt.Print(guess)

		if len(guess) != len(answer) {
			fmt.Printf("Please choose a word with %v letters\n\n", len(answer))
			continue
		}

		for i := 0; i < len(answer); i++ {

			if i >= len(guess) {
				fmt.Println("X")
				continue
			}

			if guess[i] == answer[i] {
				correct_chars++
				fmt.Print("/")
			} else {
				fmt.Print("X")
			}
		}

		if correct_chars == len(answer) {
			println("\nYou win!")
			break
		}

		attempts++
		fmt.Printf("\n%v / 3 attempts\n", attempts)
	}
	println("Game over")
}
