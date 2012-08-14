package egoistat

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type hackerNewsFeed struct {
	Hits    int
	Results []struct {
		Item struct {
			Id     int
			Points int
		}
	}
}

type HackerNewsCounter struct{}

func (c *HackerNewsCounter) urlFor(url string) string {
	return fmt.Sprintf("http://api.thriftdb.com/api.hnsearch.com/items/_search?filter[fields][url][]=%s", url)
}

func (c *HackerNewsCounter) Count(r *Request) (count int) {
	var resp *http.Response
	var dec *json.Decoder
	var err error
	var data *hackerNewsFeed
	var url = r.Url()

	if resp, err = http.Get(c.urlFor(url)); err != nil {
		return 0
	}

	dec = json.NewDecoder(resp.Body)
	if err = dec.Decode(&data); err != nil {
		return 0
	}

	if data.Hits > 0 && len(data.Results) > 0 {
		count = data.Results[0].Item.Points
	}
	return
}

func init() {
	counters["hackernews"] = new(HackerNewsCounter)
}
