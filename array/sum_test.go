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

func TestSumAll(t *testing.T) {
	t.Run("two collections", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{3, 4})
		want := []int{3, 7}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("single collection", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3})
		want := []int{6}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumTailAll(t *testing.T) {
	t.Run("two collections", func(t *testing.T) {
		got := SumTailAll([]int{1, 2}, []int{0, 4})
		want := []int{2, 4}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
