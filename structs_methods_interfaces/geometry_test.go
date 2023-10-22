package structsmethodsinterfaces

import "testing"

func TestPermieter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	if got != want {
		t.Errorf("got %2f, want %2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name string
		got  Shape
		want float64
	}{
		{name: "Rectangle", got: Rectangle{Width: 10.0, Height: 10.0}, want: 100.0},
		{name: "Circle", got: Circle{Radius: 10.0}, want: 314.1592653589793},
		{name: "Triangle", got: Triangle{Base: 10.0, Height: 10.0}, want: 50.0},
	}
	for _, tt := range areaTests {
		t.Run("Testing area of different shapes", func(t *testing.T) {
			got := tt.got.Area()
			want := tt.want
			if got != want {
				t.Errorf("got %2f, want %2f", got, want)
			}
		})
	}
}
