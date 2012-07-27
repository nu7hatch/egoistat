package egoist

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type TwitterCounter struct {}

type twitterFeed struct {
	Count int
}

func (c *TwitterCounter) urlFor(url string) string {
	return fmt.Sprintf("http://urls.api.twitter.com/1/urls/count.json?url=%s", url)
}

func (c *TwitterCounter) Count(r *Request) int {
	var resp *http.Response
	var dec *json.Decoder
	var err error
	var data twitterFeed
	
	if resp, err = http.Get(c.urlFor(r.Url())); err != nil {
		return 0
	}

	dec = json.NewDecoder(resp.Body)
	if err = dec.Decode(&data); err != nil {
		return 0
	}

	return data.Count
}

func init() {
	counters["twitter"] = new(TwitterCounter)
}