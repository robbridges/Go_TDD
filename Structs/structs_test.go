package Structs

import (
	"reflect"
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

func TestStructToMap(t *testing.T) {
	exampleTriangleMap := map[string]interface{}{"base": 3, "height": 6}
	triangle := Triangle{Base: 3, Height: 6}
	mapGot := ConvertToMap(triangle)
	if reflect.DeepEqual(mapGot, exampleTriangleMap) {
		t.Errorf("They are not similar we wanted %v, but got %v", exampleTriangleMap, mapGot)
	}
}
