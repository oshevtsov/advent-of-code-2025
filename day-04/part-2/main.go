// Day 04, part-2
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

func CountAndRemoveAccessibleRolls(input [][]byte) (int, [][]byte) {
	var count int
	var output [][]byte

	for rowIdx, row := range input {
		outputRow := make([]byte, len(row))
		copy(outputRow, row)

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
					outputRow[colIdx] = 'x'
				}
			}
		}
		output = append(output, outputRow)
	}
	return count, output
}

func main() {
	input := ParseInput("../inputs/input.txt")

	var count int
	for {
		iterCount, iterOutput := CountAndRemoveAccessibleRolls(input)

		if iterCount == 0 {
			break
		}

		count += iterCount
		input = iterOutput
	}

	log.Printf("Answer is: %d", count) // Answer is: 9083
}
