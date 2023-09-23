package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

var ans int

func StartDay4() {
	lines := readFile("day4/input.txt")

	var wg sync.WaitGroup
	calcChan := make(chan int, 5)

	// Start Workers
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go workers(i, calcChan, &wg)
	}

	for _, line := range lines {
		calcChan <- calculate(line)
	}

	close(calcChan)
	wg.Wait()

	fmt.Println(ans)
}

func workers(id int, calcChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range calcChan {
		// fmt.Printf("Worker %d received %d\n", id, n)
		ans += n
	}
}

func calculate(line []string) int {
	rangeOne := convertToInts(strings.Split(line[0], "-"))
	rangeTwo := convertToInts(strings.Split(line[1], "-"))

	if rangeInRange(rangeOne, rangeTwo) || rangeInRange(rangeTwo, rangeOne) {
		return 1
	}
	return 0
}

func convertToInts(nums []string) []int {
	var res []int
	for _, num := range nums {
		n, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		res = append(res, n)
	}
	return res
}

func rangeInRange(first []int, last []int) bool {
	// return first[0] >= last[0] && first[1] <= last[1]
	return (first[0] >= last[0] && first[0] <= last[1]) || (first[1] >= last[0] && first[1] <= last[1])
}

func readFile(name string) [][]string {
	var res [][]string
	file, err := os.Open(name)
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		spline := strings.Split(line, ",")
		res = append(res, spline)
	}

	if err = scanner.Err(); err != nil {
		log.Fatalf("There was an error scanning the file: %v\n", err)
	}
	return res
}
