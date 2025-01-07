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

type Problem struct {
	question string
	answer   string
}

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
	totalCorrectAnswers := 0
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Failed to parse CSV file.:", err)
	}
	quizProblems := parseRecords(records)
	for i, p := range quizProblems {
		userAnswer := askQuestion(i+1, p.question)
		if userAnswer == p.answer {
			totalCorrectAnswers++
		}
	}
	populateSummary(len(quizProblems), totalCorrectAnswers)
}

func parseRecords(records [][]string) []Problem {
	response := make([]Problem, len(records))
	for i, record := range records {
		response[i] = Problem{
			question: record[0],
			answer:   record[1],
		}
	}
	return response
}

func askQuestion(index int, question string) string {
	fmt.Printf("Question #%d: %s = \n", index, question)
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
