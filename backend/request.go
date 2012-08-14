package egoistat

import "time"

type Request struct {
	options map[string]string
}

func NewRequest(url string, options map[string]string) *Request {
	if options == nil {
		options = map[string]string{}
	}
	options["url"] = url
	return &Request{options: options}
}

func (r *Request) Url() string {
	return r.options["url"]
}

func (r *Request) Option(key string) (res string) {
	res, _ = r.options[key]
	return
}

func (r *Request) Stat(networks ...string) (results ResultsGroup) {
	if len(r.Url()) == 0 {
		return
	}

	results = ResultsGroup{}

	var fanin = make(chan *Result)
	var timeout = time.After(10 * time.Second)
	var jobs = 0

	for _, network := range networks {
		if counter, ok := FindCounter(network); ok {
			jobs++
			go func(network string) {
				partial := counter(r).In(network)
				fanin <- partial
			}(network)
		}
	}

	for ; jobs > 0; jobs-- {
		select {
		case partial := <-fanin:
			results.Add(partial)
		case <-timeout:
			return
		}
	}

	return
}

func (r *Request) StatAll() ResultsGroup {
	networks := []string{}
	for net, _ := range counters {
		networks = append(networks, net)
	}

	return r.Stat(networks...)
}
