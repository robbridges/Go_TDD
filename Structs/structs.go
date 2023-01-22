package Structs

import (
	"math"
	"reflect"
)

type Rectangle struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	getArea() float64
}

func ConvertToMap(item interface{}) map[string]interface{} {
	itemValue := reflect.ValueOf(item)
	itemType := itemValue.Type()

	Result := make(map[string]interface{})
	for i := 0; i < itemValue.NumField(); i++ {
		fieldValue := itemValue.Field(i)
		fieldType := itemType.Field(i)
		Result[fieldType.Name] = fieldValue.Interface()
	}

	return Result
}

func getPerimeter(rec Rectangle) float64 {
	return 2 * (rec.width + rec.length)
}

func (rec Rectangle) getArea() float64 {
	return rec.length * rec.width
}

func (cir Circle) getArea() float64 {
	return math.Pi * (cir.radius * cir.radius)
}

func (tri Triangle) getArea() float64 {
	return .5 * (tri.Base * tri.Height)
}
