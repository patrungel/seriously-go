package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func parseArgs() Opts {
	var quizFilePath = flag.String("file", "problems.csv", "Path to a CSV file with problems and answers")
	var quizTimeout = flag.Int("time", 30, "Time to take the quiz, in seconds")
	var shuffle = flag.Bool("shuffle", false, "Shuffle the problems")
	flag.Parse()

	return Opts{
		FilePath: *quizFilePath,
		Timeout:  *quizTimeout,
		Shuffle:  *shuffle,
	}
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

func runQuiz(quiz []Problem, scoreChan chan int, dismissChan chan struct{}) {
	score := 0
	for _, problem := range quiz {
		fmt.Printf("%s: ", problem.Question)
		var answer string
		_, err := fmt.Scanln(&answer)
		if err != nil {
			continue
		}
		if strings.ToLower(answer) == strings.ToLower(problem.Answer) {
			score++
			scoreChan <- score
		}
	}
	dismissChan <- struct{}{}
}

type Problem struct {
	Question string
	Answer   string
}

type Opts struct {
	FilePath string
	Timeout  int
	Shuffle  bool
}

func main() {
	opts := parseArgs()

	quizProblems, err := getQuizFromFile(opts.FilePath)
	if err != nil {
		log.Fatal(err)
	}

	if opts.Shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(quizProblems), func(i, j int) {
			quizProblems[i], quizProblems[j] = quizProblems[j], quizProblems[i]
		})
	}

	scoreChan := make(chan int, len(quizProblems))
	dismissChan := make(chan struct{})
	go runQuiz(quizProblems, scoreChan, dismissChan)

	score := 0
	timer := time.NewTimer(time.Duration(opts.Timeout) * time.Second)
	for runTimer := true; runTimer; {
		select {
		case <-dismissChan:
			runTimer = false
			timer.Stop()
		case score = <-scoreChan:
			// do nothing
		case <-timer.C:
			runTimer = false
			fmt.Println("\nTime is up.")
		}
	}
	fmt.Printf("Your score is %d out of %d.\n", score, len(quizProblems))
}
