package egoistat

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

type result struct {
	network string
	count   int
}

func (r *Request) Count(networks ...string) (res map[string]int) {
	res = make(map[string]int)
	if len(r.Url()) == 0 {
		return
	}

	counts := make(chan *result)
	defer close(counts)

	numJobs := 0
	for _, net := range networks {
		if counter, ok := counters[net]; ok {
			numJobs++
			go func(net string) {
				counts <- &result{net, counter.Count(r)}
			}(net)
		}
	}

	for i := 0; i < numJobs; i++ {
		partial := <-counts
		res[partial.network] = partial.count
	}

	return
}

func (r *Request) CountAll() map[string]int {
	networks := []string{}
	for net, _ := range counters {
		networks = append(networks, net)
	}

	return r.Count(networks...)
}
