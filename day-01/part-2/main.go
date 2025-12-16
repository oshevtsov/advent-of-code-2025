// Part 2: Main
package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ProcessDocumentPart2(path string) int {
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

		// Remove full rotations
		reminderValue := value % 100

		numZeroDials += value / 100
		prevDial := currDial
		if direction == 'L' {
			currDial -= reminderValue

			if currDial < 0 {
				// If the starting point is zero, going backwards does not count
				// as crossing zero (case of stopping at zero is covered below).
				if prevDial != 0 {
					numZeroDials++
				}
				currDial += 100
			}
		} else {
			currDial += reminderValue

			if currDial > 99 {
				// If the final point overflows 100, then we should have crossed
				// zero (case of stopping at zero is covered below).
				if currDial > 100 {
					numZeroDials++
				}
				currDial -= 100
			}
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
	numZeroDials := ProcessDocumentPart2("../inputs/input.txt")
	log.Printf("Answer is: %d", numZeroDials) // Answer is: 6295
}
