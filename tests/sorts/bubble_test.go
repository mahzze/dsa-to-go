package sorts_test

import (
	"testing"

	"github.com/mahzze/dsa-to-go/algorithms/sorts"
)

func TestBubbleSort(t *testing.T) {
	givenArrays := [][]int{
		{0, 1, 5, 3, 2, -14, 42, 7, 69, 420, 10},
		{10, 10, 10, 10, 10, 10, 10, 10, 10},
		{1, 2, 3, 4, 5, 6, 0, 7, 8, 9},
		{5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5},
	}
	wantedArrays := [][]int{
		{-14, 0, 1, 2, 3, 5, 7, 10, 42, 69, 420},
		{10, 10, 10, 10, 10, 10, 10, 10, 10},
		{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
	}
	var testArrays [][]int

	for i := range givenArrays {
		testArrays = append(testArrays, sorts.BubbleSort(givenArrays[i]))
	}
	for i := range givenArrays {
		for j := range givenArrays[i] {
			if wantedArrays[i][j] != testArrays[i][j] {
				t.Errorf("Arrays not properly sorted!\nexpected : %v\ngot : %v", wantedArrays[i], testArrays[i])
			}
		}
	}
}
