package main

import "testing"

type Shape interface {
	Area() float64
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: Rectangle{12.0, 6.0}, want: 72.0},
		{name: "circle", shape: Circle{10.0}, want: 314.1592653589793},
		{name: "triangle", shape: Triangle{12.0, 6.0}, want: 36.0},
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			want := test.want

			if got != want {
				t.Errorf("%#v got %.2f want %.2f", test.shape, got, want)
			}
		})
	}
}
