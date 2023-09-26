package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func StartDay5() {
	file, err := os.Open("day5/input.txt")
	if err != nil {
		log.Fatal("There was an error opening the file", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	boxes, boxMoves := readBoxesFromFile(file, scanner)

	rowMap := createRowMap(boxes)
	//do not send the last row with the row numbers
	rowColumns := convertBoxRowsToColumns(boxes[:len(boxes)-1], rowMap)
	retrieveMoves(rowColumns, boxMoves)

	ans := retrieveLastBox(rowColumns)
	fmt.Println(ans)

	if err = scanner.Err(); err != nil {
		log.Fatal("There was an error scanning the file", err)
	}
}

func retrieveLastBox(rows [][]string) []string {
	var ans []string
	for _, row := range rows {
		ans = append(ans, row[len(row)-1])
	}
	return ans
}

// Move the boxes given the instructions.
// Will call function to move the boxes.
func retrieveMoves(rows [][]string, moves []string) {
	for _, move := range moves {
		split := strings.Split(move, " ")
		moveCount := castToInt(split[1])
		// match zero index
		fromRow := castToInt(split[3]) - 1
		toRow := castToInt(split[5]) - 1
		moveBoxes(moveCount, fromRow, toRow, rows)
	}
}

// Actually move the boxes
func moveBoxes(count int, from int, to int, rows [][]string) {
	rowLength := len(rows[from])
	if count < rowLength {
		remainingBoxes := rows[from][:rowLength-count]
		boxesToMove := rows[from][rowLength-count:]
		rows[from] = remainingBoxes
		// for i := len(boxesToMove) - 1; i >= 0; i-- {
		// 	rows[to] = append(rows[to], boxesToMove[i])
		// }
		rows[to] = moveBoxesOfSize(boxesToMove, rows[to], 3)
	} else {
		// Take all of the 'from' row and move it to the 'to' row
		boxesToMove := rows[from]
		rows[from] = []string{}
		// for i := len(boxesToMove) - 1; i >= 0; i-- {
		// 	rows[to] = append(rows[to], boxesToMove[i])
		// }
		rows[to] = moveBoxesOfSize(boxesToMove, rows[to], 3)
	}
}

func printBoxes(boxes [][]string) {
	for idx, row := range boxes {
		fmt.Println(idx+1, row)
	}
}

func moveBoxesOfSize(boxesToMove []string, rowToMoveTo []string, sizeOfMove int) []string {
	// I thought that there was a limit to the number of boxes the crane could move at once.
	// if len(boxesToMove) <= sizeOfMove {
	//     rowToMoveTo = append(rowToMoveTo, boxesToMove...)
	// } else {
	//     for len(boxesToMove) >= sizeOfMove {
	// 	currMove := boxesToMove[len(boxesToMove)-sizeOfMove:]
	// 	boxesToMove = boxesToMove[:len(boxesToMove)-sizeOfMove]
	// 	rowToMoveTo = append(rowToMoveTo, currMove...)
	//     }
	//     if len(boxesToMove) > 0 {
	// 	rowToMoveTo = append(rowToMoveTo, boxesToMove...)
	//     }
	// }
	rowToMoveTo = append(rowToMoveTo, boxesToMove...)
	return rowToMoveTo
}

func castToInt(num string) int {
	n, err := strconv.Atoi(num)
	if err != nil {
		log.Fatal("Could not convert, NaN", err)
	}
	return n
}

// Retreives the top portion of the input that contains the box diagram.
func readBoxesFromFile(file *os.File, scanner *bufio.Scanner) ([]string, []string) {
	var boxes []string
	var moves []string
	trigger := true //true for the top half.
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			trigger = false
			continue
		}

		if trigger {
			boxes = append(boxes, line)
		} else {
			moves = append(moves, line)
		}
	}
	return boxes, moves
}

// Takes the index that the row is positioned in the string and maps it to the row number.
// For example, row 1 may start at index 5 of the string. So the result will be [5:1, ...].
func createRowMap(boxes []string) map[int]int {
	rowNumbers := boxes[len(boxes)-1]
	rowMap := make(map[int]int)

	for idx, num := range rowNumbers {
		convNum, err := strconv.Atoi(string(num))
		if err == nil {
			rowMap[idx] = convNum
		}
	}
	return rowMap
}

// Put the letters of the horizontal rows of the string to vertical rows of the row numbers.
// The index of the slice is the row number minus 1 to account for a zero index.
func convertBoxRowsToColumns(boxRows []string, rowMap map[int]int) [][]string {
	result := make([][]string, len(rowMap))
	for i := len(boxRows) - 1; i >= 0; i-- {
		for idx, letter := range boxRows[i] {
			val, exists := rowMap[idx]
			if exists && letter != ' ' {
				// minus 1 to match index zero
				result[val-1] = append(result[val-1], string(letter))
			}
		}
	}

	return result
}
