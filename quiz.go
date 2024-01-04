package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to quiz game!")

	var name string
	var age uint

	fmt.Printf("Enter your name: ")
	fmt.Scan(&name)
	fmt.Printf("Hello, %v\n", name)

	fmt.Printf("Enter your age: ")
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("You can play this game")
	} else {
		fmt.Println("You cannot play this game\nSorry")
		return
	}

	score := 0
	num_soal := 2

	fmt.Println("Continue")

	fmt.Printf("What is the capital of Indonesia? ")

	var answer string
	var answer2 string

	fmt.Scan(&answer)

	if answer == "Jakarta" || answer == "jakarta" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("The next capital of Indonesia is Nusantara or Balikpapan? ")

	fmt.Scan(&answer2)

	if answer2 == "Nusantara" || answer2 == "nusantara" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("You scored %v out of %v. \n", score, num_soal)
	percent := (float64(score) / float64(num_soal)) * 100
	fmt.Printf("%v, you scored: %v%%.", name, percent)
}