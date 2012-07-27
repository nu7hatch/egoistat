package egoist

type Counter interface {
	Count(r *Request) int
}

var counters = map[string]Counter{}