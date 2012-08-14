package egoistat

type Counter func(*Request) *Result

var counters = map[string]Counter{}

func RegisterCounter(name string, counter Counter) {
	counters[name] = counter
}

func FindCounter(name string) (c Counter, ok bool) {
	c, ok = counters[name]
	return
}
