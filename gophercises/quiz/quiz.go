package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func parseArgs() string {
	var quizFilePath = flag.String("file", "problems.csv", "Path to a CSV file with problems and answers")
	flag.Parse()
	return *quizFilePath
}

func getQuizFromFile(filename string) ([]Problem, error) {
	inputFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	quizReader := csv.NewReader(inputFile)
	quizRecords, err := quizReader.ReadAll()
	if err != nil {
		return nil, err
	}

	problems := make([]Problem, len(quizRecords))
	for i, record := range quizRecords {
		problems[i] = Problem{Question: record[0], Answer: strings.TrimSpace(record[1])}
	}
	return problems, nil
}

func runQuiz(quiz []Problem) (int, error) {
	score := 0
	for _, problem := range quiz {
		fmt.Printf("%s: ", problem.Question)
		var answer string
		_, err := fmt.Scanln(&answer)
		if err != nil {
			continue
		}
		if answer == problem.Answer {
			score++
		}
	}
	return score, nil
}

type Problem struct {
	Question string
	Answer   string
}

func main() {
	quizFilePath := parseArgs()

	quizProblems, err := getQuizFromFile(quizFilePath)
	if err != nil {
		log.Fatal(err)
	}

	score, err := runQuiz(quizProblems)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Your score is %d out of %d.\n", score, len(quizProblems))
}
