package egoistat

import "encoding/json"

type Result struct {
	Network string
	Points  int
}

func (r *Result) In(network string) *Result {
	r.Network = network
	return r
}

var Empty = &Result{}

type ResultsGroup []*Result

func (r *ResultsGroup) Add(result *Result) {
	(*r) = append((*r), result)
}

func (r ResultsGroup) Find(name string) *Result {
	for _, result := range r {
		if result.Network == name {
			return result
		}
	}
	return nil
}

func (r ResultsGroup) MarshalJSON() ([]byte, error) {
	results := make(map[string]int)
	for _, result := range r {
		results[result.Network] = result.Points
	}
	return json.Marshal(results)
}
