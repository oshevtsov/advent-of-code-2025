package main

import "testing"

func TestCountAndRemoveAccessibleRolls(t *testing.T) {
	input := ParseInput("../inputs/test.txt")

	var got int
	for {
		iterCount, iterOutput := CountAndRemoveAccessibleRolls(input)

		if iterCount == 0 {
			break
		}

		got += iterCount
		input = iterOutput
	}

	want := 43

	if want != got {
		t.Errorf("Accessible rolls count: got %d want %d", got, want)
	}
}
