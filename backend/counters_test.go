package egoist

import (
	"testing"
)

var counterTests = map[string](func(string) *Request){
	"twitter": func(url string) *Request {
		return NewRequest(url, nil)
	},
	"facebook": func(url string) *Request {
		return NewRequest(url, nil)
	},
	"plusone": func(url string) *Request {
		return NewRequest(url, nil)
	},
	"github": func(url string) *Request {
		return NewRequest(url, map[string]string{"github_repo": "nu7hatch/areyoufuckingcoding.me"})
	},
	"hackernews": func(url string) *Request {
		return NewRequest(url, nil)
	},
	"reddit": func(url string) *Request {
		return NewRequest(url, nil)
	},
}

func TestRequestCountForCounters(t *testing.T) {
	for net, test := range counterTests {
		r := test("http://github.com/")
		if count, _ := r.Count(net)[net]; count == 0 {
			t.Errorf("Expected to get count from %s, got nothing", net)
		}
	}
}