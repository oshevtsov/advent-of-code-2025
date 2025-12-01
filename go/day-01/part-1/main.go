// Part 1: Main
package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ProcessDocumentPart1(path string) int {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numZeroDials int
	currDial := 50
	for scanner.Scan() {
		currRotation := scanner.Text()
		direction := currRotation[0]
		value, err := strconv.Atoi(currRotation[1:])

		if err != nil {
			log.Fatalf("Failed to parse rotation value %s: %v", currRotation[1:], err)
		}

		if direction == 'L' {
			currDial -= value
		} else {
			currDial += value
		}

		for currDial < 0 {
			currDial += 100
		}

		for currDial > 99 {
			currDial -= 100
		}

		if currDial == 0 {
			numZeroDials++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return numZeroDials
}

func main() {
	numZeroDials := ProcessDocumentPart1("../inputs/input.txt")
	log.Printf("Answer is: %d", numZeroDials)
}
