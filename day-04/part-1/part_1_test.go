package main

import "testing"

func TestParseInput(t *testing.T) {
	input := ParseInput("../inputs/test.txt")

	numRowsGot := len(input)
	numColsGot := len(input[0])

	numRowsWant := 10
	numColsWant := 10

	if numRowsWant != numRowsGot {
		t.Errorf("Number of rows: got %d want %d", numRowsGot, numRowsWant)
	}

	if numColsWant != numColsGot {
		t.Errorf("Number of columns: got %d want %d", numColsGot, numColsWant)
	}
}

func TestCountAccessibleRolls(t *testing.T) {
	input := ParseInput("../inputs/test.txt")
	got := CountAccessibleRolls(input)
	want := 13

	if want != got {
		t.Errorf("Accessible rolls count: got %d want %d", got, want)
	}
}
