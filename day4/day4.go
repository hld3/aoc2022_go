package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func StartDay4() {
	ans := 0
	lines := readFile("day4/input.txt")
	for _, line := range lines {
		rangeOne := convertToInts(strings.Split(line[0], "-"))
		rangeTwo := convertToInts(strings.Split(line[1], "-"))

		if rangeInRange(rangeOne, rangeTwo) || rangeInRange(rangeTwo, rangeOne) {
			ans++
		}
	}
	fmt.Println(ans)
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
	//TODO determine if one range is within the other.
	return first[0] >= last[0] && first[1] <= last[1]
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
