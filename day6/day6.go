package day6

import (
	"bufio"
	"fmt"
	"os"
)

func StartDay6() {
	file, err := os.Open("day6/input.txt")
	if err != nil {
		fmt.Println("Error opening the file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		res := findMarkerIndex(line)
		fmt.Println(res)
	}
}

// negative one means no marker was found.
func findMarkerIndex(packet string) int {
	chars := []string{}
	for idx, char := range packet {
	    if len(chars) == 14 {
		if containsDupes(chars) {
		    chars = chars[1:]
		    chars = append(chars, string(char))
		} else {
		    return idx
		}
	    } else {
		chars = append(chars, string(char))
	    }
	}
	return -1
}

// loop through the slice looking for duplicates
// return true if dupes are found, false otherwise
func containsDupes(arr []string) bool {
	for i, val := range arr {
		for j := i + 1; j < len(arr); j++ {
			if val == arr[j] {
				return true
			}
		}
	}
	return false
}
