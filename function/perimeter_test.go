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
	areaTest := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{width: 12, height: 6}, 72.0},
		{Circle{radius: 10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}
	for _, tt := range areaTest {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v wanted %g but got %g", tt.shape, tt.want, got)

		}

	}

}
