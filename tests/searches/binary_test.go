package searches_test

import (
	"testing"

	"github.com/mahzze/dsa-to-go/algorithms/searches"
)

func TestBinarySearch(t *testing.T) {
	testArr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	targets := []int{11, 5, 0, 1, -1, 42, 10}
	wantedResults := [7]bool{false, true, true, true, false, false, true}

	testResults := [7]bool{}
	for k, v := range targets {
		testResults[k] = searches.BinarySearch(testArr, v)
	}

	if testResults != wantedResults {
		t.Errorf("\nWanted: %v\nGot   : %v", wantedResults, testResults)
	}
}
