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
	quizTime := flag.Int("time", 30, "Quiz timer in secs")
	flag.Parse()

	fmt.Println("Welcome to Who Wants to be a millionaire!!!")
	fmt.Printf("You have %d secs to finsih the quiz and answer all the questions!!! \n", *quizTime)
	fmt.Println("Press Enter to start your quiz!!!")
	// Prompt User to press Enter to start
	bufio.NewReader(os.Stdin).ReadString('\n')
	// Set the timer for the quiz
	timer := time.NewTimer(time.Duration(*quizTime) * time.Second)
	done := make(chan bool)

	score := 0
	var quizProblems []Problem

	go func() {
		// Open the csv file
		file, err := os.Open(*fileName)

		if err != nil {
			log.Fatal("Error opening file:", err)
		}

		defer file.Close()
		// create new csv reader
		reader := csv.NewReader(file)

		records, err := reader.ReadAll()
		if err != nil {
			log.Fatal("Failed to parse CSV file.:", err)
		}
		quizProblems = parseRecords(records)
		for i, p := range quizProblems {
			userAnswer := askQuestion(i+1, p.question)
			if userAnswer == p.answer {
				score++
			}
		}
		done <- true
	}()

	select {
	case <-timer.C:
		// Timer's expired
		fmt.Println("Time's up!")
	case <-done:
		// quiz finished before timed is complete
		fmt.Println("You completed the quiz!")
	}
	populateSummary(len(quizProblems), score)
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
