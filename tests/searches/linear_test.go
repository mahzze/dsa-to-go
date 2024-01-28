package searches_test

import (
	"testing"

	"github.com/mahzze/dsa-to-go/algorithms/searches"
)

func TestLinearSearch(t *testing.T) {
	testArr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	targets := []int{-15, 0, 10, 7, 23, 42}
	wantedResult := [6]bool{false, true, true, true, false, false}

	testResults := [6]bool{}
	for k, v := range targets {
		testResults[k] = searches.LinearSearch(testArr, v)
	}
	if testResults != wantedResult {
		t.Errorf("Error --------------\nWanted: %v\nGot: %v", wantedResult, testResults)
	}
}
