package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	expected := 40.0

	if got != expected {
		t.Errorf("got %.2f but expected %.2f", got, expected)
	}
}

func TestArea(t *testing.T) {
	t.Run("should calc rectangle area", func (t *testing.T)  {
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Area()
		expected := 100.0

		if got != expected {
			t.Errorf("got %.2f but expected %.2f", got, expected)
		}
	})

	t.Run("should calc circle area", func(t *testing.T) {
		circle :=  Circle{10.0}
		got := circle.Area()
		expected := 314.1592653589793

		if got != expected {
			t.Errorf("got %.g but expected %.g", got, expected)
		}
	})
}