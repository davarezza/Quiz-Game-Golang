package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the quiz game!")

	var name string
	var age uint

	fmt.Printf("Enter your name: ")
	fmt.Scan(&name)
	fmt.Printf("Hello, %v\n", name)

	fmt.Printf("Enter your age: ")
	fmt.Scan(&age)

	if age < 10 {
		fmt.Println("You cannot play this game\nSorry")
		return
	}

	fmt.Println("You can play this game")

	score := 0
	num_soal := 2

	fmt.Println("Continue")

	var category string
	fmt.Printf("Select the question topic (enter a number):\n1. Geography\n2. Technology\nYour choice: ")
	fmt.Scan(&category)

	switch category {
	case "1":
		fmt.Printf("1. What is the capital of Indonesia? ")
		var answer string
		fmt.Scan(&answer)

		if answer == "Jakarta" || answer == "jakarta" {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Println("Incorrect!")
		}

		fmt.Printf("2. The next capital of Indonesia is Nusantara or Balikpapan? ")
		var answer2 string
		fmt.Scan(&answer2)

		if answer2 == "Nusantara" || answer2 == "nusantara" {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Println("Incorrect!")
		}

	case "2":
		fmt.Printf("1. The popular programming languages between C or PHP is? ")
		var answer string
		fmt.Scan(&answer)

		if answer == "PHP" || answer == "php" {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Println("Incorrect!")
		}

		fmt.Printf("2. What is the most popular framework in PHP? ")
		var answer2 string
		fmt.Scan(&answer2)

		if answer2 == "Laravel" || answer2 == "laravel" {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Println("Incorrect!")
		}

	default:
		fmt.Println("Invalid category selection")
		return
	}

	fmt.Printf("You scored %v out of %v. \n", score, num_soal)
	percent := (float64(score) / float64(num_soal)) * 100
	fmt.Printf("%v, you scored: %v%%.", name, percent)
}
