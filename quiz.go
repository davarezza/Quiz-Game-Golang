package main

import (
	"fmt"
	"strings"
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

	score := playQuiz()
	num_soal := 3

	fmt.Printf("You scored %v out of %v. \n", score, num_soal)
	percent := calculatePercentage(score, num_soal)
	fmt.Printf("%v, you scored: %v%%.", name, percent)
}

func playQuiz() int {
	var score int

	fmt.Println("Continue")

	var category string
	fmt.Printf("Select the question topic (enter a number):\n1. Geography\n2. Technology\nYour choice: ")
	fmt.Scan(&category)

	switch category {
	case "1":
		score += askQuestion("What is the capital of Indonesia?", "Jakarta")
		score += askQuestion("The next capital of Indonesia is Nusantara or Balikpapan?", "Nusantara")
		score += askQuestion("How many provinces in Indonesia?", "38")
	case "2":
		score += askQuestion("The popular programming languages between C or PHP is?", "PHP")
		score += askQuestion("What is the most popular framework in PHP?", "Laravel")
		score += askQuestion("Better Intel Core i3 or Intel Core i5", "Intel Core i5")
	default:
		fmt.Println("Invalid category selection")
		return 0
	}

	return score
}

func askQuestion(question, correctAnswer string) int {
	var answer string

	fmt.Printf("%s ", question)
	fmt.Scan(&answer)

	answer = strings.TrimSpace(answer)
	correctAnswer = strings.TrimSpace(correctAnswer)

	expectedWords := strings.Fields(correctAnswer)

	found := false
	for _, word := range expectedWords {
		if strings.ToLower(answer) == strings.ToLower(word) {
			found = true
			break
		}
	}

	if found {
		fmt.Println("Correct!")
		return 1
	} else {
		fmt.Println("Incorrect!")
		return 0
	}
}

func calculatePercentage(score, totalQuestions int) float64 {
	return (float64(score) / float64(totalQuestions)) * 100
}
