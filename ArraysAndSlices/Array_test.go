package ArraysAndSlices

import (
	"reflect"
	"testing"
)

func TestGetSumFromArray(t *testing.T) {
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{60, 10, 20, 30, 40}
		want := 160
		got := GetSumFromArray(numbers)
		if got != want {
			t.Errorf("There was an error when adding the sum together it should have been %d, but %d was returned",
				want, got)
		}
	})

	t.Run("Collection of 3 numbers", func(t *testing.T) {
		numbers := []int{2, 3, 5}
		want := 10
		got := GetSumFromArray(numbers)
		if got != want {
			t.Errorf("There was an error when adding the sum together it should have been %d, but %d was returned",
				want, got)
		}
	})
}

func TestSumAllSlices(t *testing.T) {
	checkSums := func(t testing.TB, got []int, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("There was a problem combining the slices, we wanted %v, but we got %v", want, got)
		}
	}

	t.Run("Multiple slices", func(t *testing.T) {
		slices := [][]int{{1, 2}, {3, 4}, {5, 6}}
		want := []int{3, 7, 11}
		got := SumAllSlices(slices)
		checkSums(t, got, want)
	})
	t.Run("Single slice", func(t *testing.T) {
		slices := [][]int{{1, 2}}
		want := []int{3}
		got := SumAllSlices(slices)
		checkSums(t, got, want)
	})
	t.Run("First slice empty", func(t *testing.T) {
		slices := [][]int{{}, {1, 2}}
		want := []int{0, 3}
		got := SumAllSlices(slices)
		checkSums(t, got, want)
	})
}
