package ints

import "testing"

var testArray = []int{5, 6, 4, 3, 2}

func Test_adder(t *testing.T) {
	want := 8
	got := AddSum(4, 4)
	if got != want {
		t.Errorf("Incorrect value, wanted %d, got %d", want, got)
	}
}

func Test_slice_adder(t *testing.T) {
	want := 20
	got := AddAllNums(testArray)
	if got != want {
		t.Errorf("We did not add all the numbers correctly, we wanted %d, but got %d", want, got)
	}
}
