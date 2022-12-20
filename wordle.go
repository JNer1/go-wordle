package main

import "fmt"

func main() {
	answer := "hello"
	attempts := 0
	correct_chars := 0

	var text_from_input string

	for attempts < 3 {
		fmt.Scan(&text_from_input)

		if len(text_from_input) != len(answer) {
			fmt.Printf("Please choose a word with %v letters\n\n", len(answer))
			continue
		}

		for i := 0; i < len(answer); i++ {

			if i >= len(text_from_input) {
				fmt.Println("X")
				continue
			}

			if text_from_input[i] == answer[i] {
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
