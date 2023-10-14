package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Should return sum equals 15", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15

		sumAssertionChecker(t, got, expected, numbers)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Should sum all and retrieve a slice with results", func(t *testing.T) {
		got := SumAll([]int{1 ,2}, []int{0, 9})
		expected := []int{3, 9}

		sumAllAssertionChecker(t, got, expected)
	})
}

func sumAssertionChecker(t *testing.T, got int, expected int, numbers []int) {
	if got != expected {
		t.Helper()
		t.Errorf("got %d, but expected %d given %v", got, expected, numbers)
	}
}

func sumAllAssertionChecker(t *testing.T, got []int, expected []int) {
	if !reflect.DeepEqual(got, expected) {
		t.Helper()
		t.Errorf("got %v, but expected %v", got, expected)
	}
}