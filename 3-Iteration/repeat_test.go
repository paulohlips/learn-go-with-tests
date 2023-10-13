package iteration

import "testing"

func TestRepeat(t *testing.T) {
	result := Repeat("a")
	expected := "aaaaa"
	repeatAssertion(t, result, expected)

}

func repeatAssertion(t *testing.T, result string, expected string) {
	if result != expected {
		t.Errorf("expected %q but result %q", expected, result)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}