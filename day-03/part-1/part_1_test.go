package main

import (
	"fmt"
	"testing"
)

func TestFindMaxBankJoltage(t *testing.T) {
	tests := []struct {
		bank string
		want int
	}{
		{"987654321111111", 98},
		{"811111111111119", 89},
		{"234234234234278", 78},
		{"818181911112111", 92},
	}

	var totalWant int
	var totalGot int
	for _, test := range tests {
		t.Run(fmt.Sprintf("Bank %s", test.bank), func(t *testing.T) {
			want := test.want
			got := FindMaxBankJoltage(test.bank)

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}

			totalGot += got
			totalWant += want
		})
	}

	if totalGot != totalWant {

		t.Errorf("totalGot %d totalWant %d", totalGot, totalWant)
	}
}
