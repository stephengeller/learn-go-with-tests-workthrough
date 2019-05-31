package internal

import "testing"

func TestPerimeter(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	assertResult := func(t *testing.T, got, want float64) {
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("first test", func(t *testing.T) {
		rectangle := Rectangle{10, 10}
		got := rectangle.Perimeter()
		want := 40.0

		assertResult(t, got, want)
	})

	t.Run("area", func(t *testing.T) {
		rectangle := Rectangle{10, 10}
		want := 100.0

		checkArea(t, rectangle, want)
	})

	t.Run("circle area", func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793

		checkArea(t, circle, want)
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{12, 6}, hasArea: 72.0},
		{shape: Circle{10}, hasArea: 314.1592653589793},
		{shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasArea)
		}
	}
}
