package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	_ = iota
	ROCK
	PAPER
	SCISSORS
)

func Start() {

	fileData, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fileData.Close()

	score := 0
	scanner := bufio.NewScanner(fileData)
	for scanner.Scan() {
		res := strings.Split(scanner.Text(), " ")
		player := retrieveHandScore(res[0])
		me := retrieveHandScoreMe(player, res[1])
		score += me + retrieveResultScore(player, me)
	}
	fmt.Println(score)
}

func retrieveHandScore(play string) int {
	switch play {
	case "A":
		return ROCK
	case "B":
		return PAPER
	case "C":
		return SCISSORS
	default:
		return 0
	}
}

func retrieveHandScoreMe(player int, me string) int {
	if me == "X" {
		if player == ROCK {
			return SCISSORS
		}
		if player == PAPER {
			return ROCK
		}
		if player == SCISSORS {
			return PAPER
		}
	}
	if me == "Z" {
		if player == ROCK {
			return PAPER
		}
		if player == PAPER {
			return SCISSORS
		}
		if player == SCISSORS {
			return ROCK
		}
	}
	return player
}

func retrieveResultScore(player int, me int) int {
	switch {
	case (player == ROCK && me == PAPER) ||
		(player == PAPER && me == SCISSORS) ||
		(player == SCISSORS && me == ROCK):
		return 6
	case player == me:
		return 3
	default:
		return 0

	}
}
