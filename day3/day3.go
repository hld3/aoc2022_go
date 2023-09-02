package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Declare itemPriorities as a global variable
var itemPriorities = make(map[rune]int)

// Initialize the itemPriorities map
func init() {
	// Assign priorities for lowercase item types 'a' through 'z'
	for i, ch := 0, 'a'; ch <= 'z'; i, ch = i+1, ch+1 {
		itemPriorities[ch] = i + 1
	}

	// Assign priorities for uppercase item types 'A' through 'Z'
	for i, ch := 0, 'A'; ch <= 'Z'; i, ch = i+1, ch+1 {
		itemPriorities[ch] = i + 27
	}
}

func StartDay3() {
	ans := 0

	fileData, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatalf("There was an error opening the file: %v\n", err)
	}
	defer fileData.Close()

	var lineCount int
	var lines []string
	scanner := bufio.NewScanner(fileData)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		lineCount++

		if lineCount == 3 {
			sticker := compareGroups(lines)
			ans += tallyPoints(sticker)
			lines = []string{}
			lineCount = 0
		}
		// mid := len(line) / 2
		// first, second := line[:mid], line[mid:]
		// dupes := compareHalves(first, second)
		// ans += tallyPoints(dupes)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("There was an error scanning the file: %v\n", err)
	}
	fmt.Println(ans)
}

func tallyPoints(dupes []rune) int {
	var ans int
	for _, char := range dupes {
		ans += itemPriorities[char]
	}
	return ans
}

func compareGroups(group []string) []rune {
	var ans []rune
	fullSet := initialSet(group[0])
	matchSetWithOne := findMatchingRunes(fullSet, group[1])
	finalSet := findMatchingRunes(matchSetWithOne, group[2])

	for char := range finalSet {
		ans = append(ans, char)
	}
	return ans
}

func initialSet(s string) map[rune]bool {
	set := make(map[rune]bool)
	// find distinct values in the first string.
	for _, char := range s {
		set[char] = true
	}
	return set
}

func findMatchingRunes(set map[rune]bool, s string) map[rune]bool {
	sSet := make(map[rune]bool)
	for _, char := range s {
		if _, exists := set[char]; exists {
			sSet[char] = true
		}
	}
	return sSet
}

func compareHalves(first string, last string) []rune {
	var ans []rune
	temp := make(map[rune]bool)
	set := make(map[rune]bool)
	// find distinct values in the first string.
	for _, char := range first {
		set[char] = true
	}
	// find distinct values in the second string that are also in the first
	for _, char := range last {
		if _, exists := set[char]; exists {
			temp[char] = true
		}
	}
	// put distinct values that are in both into a slice
	for char := range temp {
		ans = append(ans, char)
	}
	return ans
}
