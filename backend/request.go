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

type result struct {
	network string
	count   int
}

func (r *Request) Count(networks ...string) (res map[string]int) {
	res = make(map[string]int)
	if len(r.Url()) == 0 {
		return
	}

	var results = make(chan *result)
	var quit = make(chan bool)
	var numJobs = 0
	var done = false
	var timeout = time.After(10 * time.Second)
	
	defer func() {
		close(results)
		close(quit)
	}()

	for _, net := range networks {
		if counter, ok := counters[net]; ok {
			numJobs++
			go func(net string, quit <-chan bool) {
				select {
				case results <- &result{net, counter.Count(r)}:
				case <-quit:
				}
			}(net, quit)
		}
	}

	for i := 0; i < numJobs; i++ {
		if done {
			quit <- true
			continue
		}
		select {
		case partial := <-results:
			res[partial.network] = partial.count
		case <-timeout:
			done = true
		}
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
