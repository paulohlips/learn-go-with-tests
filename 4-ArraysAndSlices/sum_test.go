package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	expected := 15

	if got != expected {
		t.Errorf("got %d, but expected %d given %v", got, expected, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1 ,2}, []int{0, 9})
	expected := []int{3, 9}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v, but expected %v", got, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v, but expected %v", got, expected)
		}
	}
	t.Run("Should sum slice tail", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		checkSums(t, got, expected)
	})

	t.Run("Should sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3})
		expected := []int{0, 5}

		checkSums(t, got, expected)
	})
}