package egoistat

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const twitterFeedUrl = "http://urls.api.twitter.com/1/urls/count.json?url=%s"

type twitterFeed struct {
	Count int
}

func CountTweets(r *Request) *Result {
	var resp *http.Response
	var dec *json.Decoder
	var err error
	var feed twitterFeed
	var feedUrl = fmt.Sprintf(twitterFeedUrl, r.Url())

	if resp, err = http.Get(feedUrl); err != nil {
		return Empty
	}

	dec = json.NewDecoder(resp.Body)
	if err = dec.Decode(&feed); err != nil {
		return Empty
	}

	return &Result{Points: feed.Count}
}

func init() {
	RegisterCounter("twitter", CountTweets)
}
