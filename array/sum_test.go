package main

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := sum(numbers)
	want := 15

	if want != got {
		t.Errorf("got %d and want %d give, %v", got, want, numbers)
	}
}
