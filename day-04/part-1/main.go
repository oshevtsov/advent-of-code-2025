// Day 04, part-1
package main

import (
	"bufio"
	"log"
	"os"
)

func ParseInput(path string) [][]byte {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var input [][]byte
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		input = append(input, []byte(line))
	}

	return input
}

func CountAccessibleRolls(input [][]byte) int {
	var count int
	for rowIdx, row := range input {
		rowLen := len(row)
		for colIdx, c := range row {
			if c == '@' {
				var numNeighbors int

				// check same row
				if colIdx > 0 && row[colIdx-1] == '@' {
					numNeighbors++
				}

				if colIdx < rowLen-1 && row[colIdx+1] == '@' {
					numNeighbors++
				}

				// check above
				if rowIdx > 0 {
					rowAbove := input[rowIdx-1]
					if rowAbove[colIdx] == '@' {
						numNeighbors++
					}

					if colIdx > 0 && rowAbove[colIdx-1] == '@' {
						numNeighbors++
					}

					if colIdx < rowLen-1 && rowAbove[colIdx+1] == '@' {
						numNeighbors++
					}
				}

				// check below
				if rowIdx < rowLen-1 {
					rowBelow := input[rowIdx+1]
					if rowBelow[colIdx] == '@' {
						numNeighbors++
					}

					if colIdx > 0 && rowBelow[colIdx-1] == '@' {
						numNeighbors++
					}

					if colIdx < rowLen-1 && rowBelow[colIdx+1] == '@' {
						numNeighbors++
					}
				}

				if numNeighbors < 4 {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	input := ParseInput("../inputs/input.txt")

	count := CountAccessibleRolls(input)
	log.Printf("Answer is: %d", count) // Answer is: 1502
}
