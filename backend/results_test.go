package egoistat

import "testing"

func TestResultsGroupAddAndFind(t *testing.T) {
	results := ResultsGroup{}
	results.Add(&Result{Points: 1, Network: "test"})
	if result := results.Find("test"); result == nil {
		t.Errorf("Expected to set and find proper result, got nothing")
	}
}
