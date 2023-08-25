package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Start() int {
	fileInput, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fileInput.Close()

	ans := make([]int, 3)
	var sumPer int

	scanner := bufio.NewScanner(fileInput)
	for scanner.Scan() {
		if scanner.Text() != "" {
			res, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			sumPer += res
		} else {
			maxIn(ans, sumPer)
			sumPer = 0
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return sumSlice(ans)
}

func sumSlice(a []int) int {
    ans := 0
    for _, num := range a {
	ans += num
    }
    return ans
}

func maxIn(a []int, b int) {
	var hold int
	hold = b
	for i := 0; i < 3; i++ {
		if hold > a[i] {
			temp := a[i]
			a[i] = hold
			hold = temp
		}
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
