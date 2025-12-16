// Part 1
package main

import (
	"bufio"
	"log"
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

func FindMaxBankJoltage(bank string) int {
	l := len(bank)

	firstDigitIdx, firstDigit := FindMaxDigit(bank[:l-1])
	_, secondDigit := FindMaxDigit(bank[firstDigitIdx+1:])

	return firstDigit*10 + secondDigit
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
		result += FindMaxBankJoltage(bank)
	}

	log.Printf("Answer is: %d", result) // Answer is: 17346
}
