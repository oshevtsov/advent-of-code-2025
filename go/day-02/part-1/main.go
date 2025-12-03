// Main: Part 1
package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ProcessRange(start, stop int) int {
	log.Printf("Start: %d and Stop: %d", start, stop)
	return 0
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

	stop, err := strconv.Atoi(bounds[0])

	if err != nil {
		log.Fatal(err)
	}

	return start, stop
}

func ProcessInput(input string) int {
	ranges := strings.Split(input, ",")

	var result int
	for _, currRange := range ranges {
		start, stop := ParseRange(currRange)
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
