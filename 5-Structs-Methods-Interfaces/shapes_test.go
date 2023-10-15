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
	checkArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		got := shape.Area()
		if got != expected {
			t.Errorf("got %.g but expected %.g", got, expected)
		}
	}
	t.Run("should calc rectangle area", func (t *testing.T)  {
		rectangle := Rectangle{10.0, 10.0}
		checkArea(t, rectangle, 100.0)
	})

	t.Run("should calc circle area", func(t *testing.T) {
		circle :=  Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}

func TestAreaTableTest(t *testing.T) {
	areaTests := []struct {
		name string
		shape Shape
		hasArea  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12,Height: 6}, hasArea:72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}