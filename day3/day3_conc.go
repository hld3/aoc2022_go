package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

// Declare itemPriorities_conc as a global variable
var itemPriorities_conc = make(map[rune]int)

// Initialize the itemPriorities_conc map
func init() {
	// Assign priorities for lowercase item types 'a' through 'z'
	for i, ch := 0, 'a'; ch <= 'z'; i, ch = i+1, ch+1 {
		itemPriorities_conc[ch] = i + 1
	}

	// Assign priorities for uppercase item types 'A' through 'Z'
	for i, ch := 0, 'A'; ch <= 'Z'; i, ch = i+1, ch+1 {
		itemPriorities_conc[ch] = i + 27
	}
}

func StartDay3_conc() {
	fmt.Println("Running concurrently")
	ans := 0

	fileData, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatalf("There was an error opening the file: %v\n", err)
	}
	defer fileData.Close()

	allLines := groupAllLines(fileData)
	// create waitgroup and chan
	var wg sync.WaitGroup
	totalsChan := make(chan int, len(allLines))

	for _, lines := range allLines {
	    wg.Add(1)
	    go processGroups(lines, &wg, totalsChan)
	}

	func () {
	    wg.Wait()
	    close(totalsChan)
	}()

	for c := range totalsChan {
	    ans += c
	}

	fmt.Println(ans)
}

func processGroups(lines []string, wg *sync.WaitGroup, c chan int) {
    defer wg.Done()

    res := compareGroups(lines)
    c <- tallyPoints(res)
}

func groupAllLines(fileData *os.File) [][]string {
	var lineCount int
	var groupLines []string
	var allLines [][]string
	scanner := bufio.NewScanner(fileData)
	for scanner.Scan() {
		line := scanner.Text()
		groupLines = append(groupLines, line)
		lineCount++

		if lineCount == 3 {
			allLines = append(allLines, groupLines)
			groupLines = []string{}
			lineCount = 0
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("There was an error scanning the file: %v\n", err)
	}
	return allLines
}

func tallyPoints_conc(dupes []rune) int {
	var ans int
	for _, char := range dupes {
		ans += itemPriorities_conc[char]
	}
	return ans
}

func compareGroups_conc(group []string) []rune {
	var ans []rune
	fullSet := initialSet_conc(group[0])
	matchSetWithOne := findMatchingRunes_conc(fullSet, group[1])
	finalSet := findMatchingRunes_conc(matchSetWithOne, group[2])

	for char := range finalSet {
		ans = append(ans, char)
	}
	return ans
}

func initialSet_conc(s string) map[rune]bool {
	set := make(map[rune]bool)
	// find distinct values in the first string.
	for _, char := range s {
		set[char] = true
	}
	return set
}

func findMatchingRunes_conc(set map[rune]bool, s string) map[rune]bool {
	sSet := make(map[rune]bool)
	for _, char := range s {
		if _, exists := set[char]; exists {
			sSet[char] = true
		}
	}
	return sSet
}

func compareHalves_conc(first string, last string) []rune {
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
