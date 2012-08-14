package egoistat

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const hackerNewsFeedUrl = "http://api.thriftdb.com/api.hnsearch.com/items/_search?filter[fields][url][]=%s"

type hackerNewsFeed struct {
	Hits    int
	Results []struct {
		Item struct {
			Id     int
			Points int
		}
	}
}

func CountHackerNewsPoints(r *Request) *Result {
	var resp *http.Response
	var dec *json.Decoder
	var err error
	var feed hackerNewsFeed
	var feedUrl = fmt.Sprintf(hackerNewsFeedUrl, r.Url())

	if resp, err = http.Get(feedUrl); err != nil {
		return Empty
	}

	dec = json.NewDecoder(resp.Body)
	if err = dec.Decode(&feed); err != nil {
		return Empty
	}
	if feed.Hits == 0 || len(feed.Results) == 0 {
		return Empty
	}

	return &Result{Points: feed.Results[0].Item.Points}
}

func init() {
	RegisterCounter("hackernews", CountHackerNewsPoints)
}
