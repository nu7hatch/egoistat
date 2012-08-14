package egoistat

import "testing"

func testCounter(r *Request) *Result {
	return &Result{Points: 10}
}

func init() {
	RegisterCounter("test", testCounter)
}

func TestNewRequestWithoutOptions(t *testing.T) {
	r := NewRequest("http://nu7hat.ch/", nil)
	if r.Url() != "http://nu7hat.ch/" {
		t.Errorf("Expected to set proper url, got %s", r.Url())
	}
}

func TestNewRequestWithOptions(t *testing.T) {
	r := NewRequest("http://nu7hat.ch/", map[string]string{"foo": "bar"})
	if r.Url() != "http://nu7hat.ch/" || r.Option("foo") != "bar" {
		t.Errorf("Expected to set proper options, got %v", r.options)
	}
}

func TestRequestCount(t *testing.T) {
	r := NewRequest("http://nu7hat.ch/", nil)
	if result := r.Stat("test").Find("test"); result.Points != 10 {
		t.Errorf("Expected to get proper count, got %v", result.Points)
	}
}
