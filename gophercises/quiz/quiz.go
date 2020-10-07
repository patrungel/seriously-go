package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func parseArgs() string {
	var quizFilePath = flag.String("file", "problems.csv", "Path to a CSV file with problems and answers")
	flag.Parse()
	return *quizFilePath
}

func getQuizFromFile(filename string) ([][]string, error) {
	inputFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	quizReader := csv.NewReader(inputFile)
	quizRecords, err := quizReader.ReadAll()
	if err != nil {
		return  nil, err
	}

	return quizRecords, nil
}

func runQuiz(quiz [][]string) (int, error) {
	score := 0
	for _, record := range quiz {
		fmt.Printf("%s: ", record[0])
		var answer string
		_, err := fmt.Scanln(&answer)
		if err != nil {
			continue
		}
		if answer == record[1] {
			score++
		}
	}
	return score, nil
}

func main() {
	quizFilePath := parseArgs()

	quizRecords, err := getQuizFromFile(quizFilePath)
	if err != nil {
		log.Fatal(err)
	}

	score, err := runQuiz(quizRecords)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Your score is %d out of %d.\n", score, len(quizRecords))
}
