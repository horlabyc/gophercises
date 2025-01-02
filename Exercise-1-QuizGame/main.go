package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	// get CSV file name
	fileName := flag.String("file_name", "problems.csv", "Path to the CSV file")
	flag.Parse()

	// Open the csv file
	file, err := os.Open(*fileName)

	if err != nil {
		log.Fatal("Error opening file:", err)
	}

	defer file.Close()
	fmt.Println("Quiz loading, pls wait...")
	time.Sleep(3 * time.Second)
	fmt.Println("Enter your answers after the questions...")
	// create new csv reader
	reader := csv.NewReader(file)
	totalQuestions := 0
	totalCorrectAnswers := 0
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			continue
		}
		question := record[0]
		answer := record[1]
		userAnswer := askQuestion(question)
		if userAnswer == answer {
			totalCorrectAnswers++
		}
		totalQuestions++
	}
	populateSummary(totalQuestions, totalCorrectAnswers)
}

func askQuestion(question string) string {
	fmt.Printf("Question: %-10v\n", question)
	inputReader := bufio.NewReader(os.Stdin)
	opt, _ := inputReader.ReadString('\n')
	return strings.TrimSpace(opt)
}

func populateSummary(totalQuestions int, totalCorrectAnswers int) {
	score := (float64(totalCorrectAnswers) / float64(totalQuestions)) * 100
	fmt.Println("=============================================================================")
	fmt.Printf("Total number of questions: %-10v\n", totalQuestions)
	fmt.Printf("Total number of questions answered correctly: %-10v\n", totalCorrectAnswers)
	fmt.Printf("Quiz Score: %0.2f%%\n", score)
}
