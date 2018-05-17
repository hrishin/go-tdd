package main

import "testing"
import "reflect"

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func assertSlices(t *testing.T, got, want []int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAll(t *testing.T) {
	t.Run("two collections", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{3, 4})
		want := []int{3, 7}

		assertSlices(t, got, want)
	})

	t.Run("single collection", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3})
		want := []int{6}

		assertSlices(t, got, want)
	})
}

func TestSumTailAll(t *testing.T) {

	t.Run("two collections", func(t *testing.T) {
		got := SumTailAll([]int{1, 2}, []int{0, 4})
		want := []int{2, 4}

		assertSlices(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumTailAll([]int{}, []int{2, 3})
		want := []int{0, 3}

		assertSlices(t, got, want)
	})
}
