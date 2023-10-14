package arraysandslices

import "testing"

func TestSum(t *testing.T) {
	t.Run("Should return sum equals 15", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15

		sumAssertionChecker(t, got, expected, numbers)
	})
}

func sumAssertionChecker(t *testing.T, got int, expected int, numbers []int) {
	if got != expected {
		t.Errorf("got %d, but expected %d given %v", got, expected, numbers)
	}
}