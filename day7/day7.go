package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"strings"
)

type DirectoryData struct {
	Id   int
	Size int
}

func StartDay7() {
	file, err := os.Open("day7/input.txt")
	if err != nil {
		log.Fatal("Error opening file", err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	dirSize := retrieveDirSizes(lines)
	ans := retrieveSizeToRemove(dirSize)
	fmt.Println(ans)
}

func retrieveSizeToRemove(dirSize []DirectoryData) int {
	res := 0
	for _, dir := range dirSize {
		if dir.Size <= 100000 {
			res += dir.Size
		}
	}
	return res
}

func retrieveDirSizes(lines []string) []DirectoryData {

	dirTracking := []int{}
	dirSize := []DirectoryData{}

	dirId := 0
	for _, line := range lines {
		lsplit := strings.Split(line, " ")

		dd := DirectoryData{}
		dd.Id = dirId

		if lsplit[1] == "cd" {
			if lsplit[2] == ".." {
				// when "cd" is "..", remove the last directory from tracking slice
				dirTracking = dirTracking[:len(dirTracking)-1]
			} else {
				// when "cd" is not "..", add the directory to the list
				dirTracking = append(dirTracking, dirId)
				dirSize = append(dirSize, dd)
				dirId++
			}
			continue
		}
		if num, isNum := convertToInt(lsplit[0]); isNum {
			for i := 0; i < len(dirTracking); i++ {
				// I am going to try to match the number in the tracking slice with the number in the size slice
				// in order to only update the dir size of the slice that matches between them.
				for idx, dir := range dirSize {
					if dir.Id == dirTracking[i] {
						dirSize[idx].Size += num
					}
				}
			}
		}
	}
	return dirSize
}

func convertToInt(num string) (int, bool) {
	val, err := strconv.Atoi(num)
	if err != nil {
		return -1, false
	}
	return val, true
}
