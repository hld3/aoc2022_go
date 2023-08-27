package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

const (
	_ = iota
	ROCK2
	PAPER2
	SCISSORS2
)

func Start2() {

	fileData, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fileData.Close()

	fmt.Println(retrieveTotalScore(fileData))
}

func retrieveTotalScore(fileData *os.File) int {
	var wg sync.WaitGroup
	var total int

	// Read all lines into a slice.
	var lines []string
	scanner := bufio.NewScanner(fileData)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check if there was an error during scanning the file.
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	totalChan := make(chan int, len(lines))
	// Process each line concurrently.
	for _, line := range lines {
		wg.Add(1)
		go retrieveLineScore(line, totalChan, &wg)
	}

	go func() {
		wg.Wait()
		close(totalChan)
	}()

	for result := range totalChan {
		total += result
	}
	return total
}

func retrieveLineScore(line string, tChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	res := strings.Split(line, " ")
	playerChoice := retrieveOppHandResult(res[0])
	myChoice := retrieveMyHandResult(playerChoice, res[1])
	tChan <- myChoice + retrieveResultScore2(playerChoice, myChoice)
}

func retrieveOppHandResult(play string) int {
	switch play {
	case "A":
		return ROCK2
	case "B":
		return PAPER2
	case "C":
		return SCISSORS2
	default:
		return 0
	}
}

func retrieveMyHandResult(player int, me string) int {
	if me == "X" {
		if player == ROCK2 {
			return SCISSORS2
		}
		if player == PAPER2 {
			return ROCK2
		}
		if player == SCISSORS2 {
			return PAPER2
		}
	}
	if me == "Z" {
		if player == ROCK2 {
			return PAPER2
		}
		if player == PAPER2 {
			return SCISSORS2
		}
		if player == SCISSORS2 {
			return ROCK2
		}
	}
	return player
}

func retrieveResultScore2(player int, me int) int {
	switch {
	case (player == ROCK2 && me == PAPER2) ||
		(player == PAPER2 && me == SCISSORS2) ||
		(player == SCISSORS2 && me == ROCK2):
		return 6
	case player == me:
		return 3
	default:
		return 0

	}
}
