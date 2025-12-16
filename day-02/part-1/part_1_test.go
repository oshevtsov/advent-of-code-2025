package main

import (
	"fmt"
	"os"
	"testing"
)

func TestParseRange(t *testing.T) {
	tests := []struct {
		start      int
		stop       int
		invalidIDs []int
	}{
		{11, 22, []int{11, 22}},
		{95, 115, []int{99}},
		{998, 1012, []int{1010}},
		{1188511880, 1188511890, []int{1188511885}},
		{222220, 222224, []int{222222}},
		{1698522, 1698528, []int{}},
		{446443, 446449, []int{446446}},
		{38593856, 38593862, []int{38593859}},
		{565653, 565659, []int{}},
		{824824821, 824824827, []int{}},
		{2121212118, 2121212124, []int{}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Range %d-%d", test.start, test.stop), func(t *testing.T) {
			got := ProcessRange(test.start, test.stop)
			want := sum(test.invalidIDs)

			if got != want {
				t.Errorf("got %d want %d", got, want)
			}
		})
	}
}

func TestProcessInput(t *testing.T) {
	content, err := os.ReadFile("../inputs/test.txt")

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	got := ProcessInput(string(content))
	want := 1227775554

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func sum(invalidIDs []int) int {
	var sum int
	for _, val := range invalidIDs {
		sum += val
	}
	return sum
}
