// Part 2
package main

import (
	"bufio"
	"log"
	"math"
	"os"
)

func FindMaxDigit(bank string) (int, int) {
	var maxIdx int
	var maxDigit int

	for idx, r := range bank {
		// convert digit rune to int
		digit := int(r - '0')
		if digit > maxDigit {
			maxDigit = digit
			maxIdx = idx
		}
	}

	return maxIdx, maxDigit
}

func FindMaxBankJoltage(bank string, numBatteries int) int {
	l := len(bank)

	var maxJoltage int
	var startIdx int

	factor := int(math.Pow10(numBatteries))
	for i := 1; i <= numBatteries; i++ {
		digitIdx, digit := FindMaxDigit(bank[startIdx : l-(numBatteries-i)])
		factor /= 10
		maxJoltage += digit * factor
		startIdx += digitIdx + 1
	}

	return maxJoltage
}

func main() {
	file, err := os.Open("../inputs/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var result int
	for scanner.Scan() {
		bank := scanner.Text()
		result += FindMaxBankJoltage(bank, 12)
	}

	log.Printf("Answer is: %d", result) // Answer is: 172981362045136
}
