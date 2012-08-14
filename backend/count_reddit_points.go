package egoistat

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const redditFeedUrl = "http://www.reddit.com/api/info.json?url=%s"

type redditFeed struct {
	Data struct {
		Children []struct {
			Data struct {
				Id        string
				Permalink string
				Score     int
			}
		}
	}
}

func CountRedditPoints(r *Request) *Result {
	var resp *http.Response
	var dec *json.Decoder
	var err error
	var feed redditFeed
	var feedUrl = fmt.Sprintf(redditFeedUrl, r.Url())

	if resp, err = http.Get(feedUrl); err != nil {
		return Empty
	}

	dec = json.NewDecoder(resp.Body)
	if err = dec.Decode(&feed); err != nil {
		return Empty
	}
	if len(feed.Data.Children) == 0 {
		return Empty
	}
	return &Result{Points: feed.Data.Children[0].Data.Score}
}

func init() {
	RegisterCounter("reddit", CountRedditPoints)
}
