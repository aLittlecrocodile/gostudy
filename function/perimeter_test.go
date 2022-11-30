package function

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{height: 10.0, width: 10.0}
	got := Perimeter(rectangle)
	wanted := 40.0
	if got != wanted {
		t.Errorf("wanted %.2f but got %.2f", wanted, got)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("want %g but got %g", want, got)
		}

	}
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{height: 6.0, width: 7.0}
		checkArea(t, rectangle, 42.0)

	})
	t.Run("circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)

	})

}
