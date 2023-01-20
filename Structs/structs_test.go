package Structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 10}
	want := 40.0
	got := getPerimeter(rectangle)
	if got != want {
		t.Errorf("We aren't calculating the area right, it should be %.2f, but got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {

	//checkArea := func(t testing.TB, shape Shape, hasArea float64) {
	//	t.Helper()
	//	got := shape.getArea()
	//	if got != hasArea {
	//		t.Errorf("We aren't calculating the area right, it should be %g, but got %g", hasArea, got)
	//	}
	//}
	//
	//t.Run("Rectangle area", func(t *testing.T) {
	//	rectangle := Rectangle{12.2, 13.4}
	//	checkArea(t, rectangle, 163.48)
	//})
	//t.Run("Circle area", func(t *testing.T) {
	//	circle := Circle{10}
	//	checkArea(t, circle, 314.1592653589793)
	//})

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.getArea()
		if got != tt.hasArea {
			t.Errorf("%#v got %g hasArea %g", tt.shape, got, tt.hasArea)
		}
	}
}
