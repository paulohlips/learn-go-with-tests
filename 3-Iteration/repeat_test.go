package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Should repeat char a 5 times", func(t *testing.T) {
		result := Repeat("a", 5)
		expected := "aaaaa"
		repeatAssertion(t, result, expected)
	})

	t.Run("Should repeat x 10 times", func(t *testing.T) {
		result := Repeat("x", 10)
		expected := "xxxxxxxxxx"
		repeatAssertion(t, result, expected)
	})

}

func repeatAssertion(t *testing.T, result string, expected string) {
	if result != expected {
		t.Errorf("expected %q but result %q", expected, result)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", b.N)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 5)
	fmt.Println(result)
	// Output: aaaaa
}
