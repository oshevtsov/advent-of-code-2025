// Main: Part 2
package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ProcessRange(start, stop int) int {
	currVal := start

	var rangeResult int
	for currVal <= stop {
		currValStr := strconv.Itoa(currVal)

		length := len(currValStr)
		chunkSize := 1

		for chunkSize <= length/2 {
			if length%chunkSize == 0 {
				firstChunk := currValStr[:chunkSize]

				nextChunkIdx := 1
				matches := true
				for {
					nextChunkStart := chunkSize * nextChunkIdx
					nextChunkEnd := chunkSize * (nextChunkIdx + 1)

					if nextChunkStart >= length {
						break
					}

					nextChunk := currValStr[nextChunkStart:nextChunkEnd]
					if nextChunk != firstChunk {
						matches = false
						break
					}

					nextChunkIdx++
				}

				// if the ID qualifies as invalid, skip increasing chunk size and
				// proceed to the next ID
				if matches {
					rangeResult += currVal
					break
				}
			}

			chunkSize++
		}

		currVal++
	}
	return rangeResult
}

func ParseRange(currRange string) (int, int) {
	bounds := strings.Split(currRange, "-")

	if len(bounds) != 2 {
		log.Fatalf("Invalid range %s, expected two values separated by a '-'", currRange)
	}

	start, err := strconv.Atoi(bounds[0])

	if err != nil {
		log.Fatal(err)
	}

	stop, err := strconv.Atoi(bounds[1])

	if err != nil {
		log.Fatal(err)
	}

	return start, stop
}

func ProcessInput(input string) int {
	ranges := strings.Split(input, ",")

	var result int
	for _, currRange := range ranges {
		start, stop := ParseRange(strings.TrimSpace(currRange))
		result += ProcessRange(start, stop)
	}

	return result
}

func main() {
	content, err := os.ReadFile("../inputs/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	input := string(content)
	result := ProcessInput(input)

	log.Printf("Answer is: %d\n", result)
}
