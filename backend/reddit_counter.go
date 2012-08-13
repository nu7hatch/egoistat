package egoist

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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

type RedditCounter struct{}

func (c *RedditCounter) urlFor(url string) string {
	return fmt.Sprintf("http://www.reddit.com/api/info.json?url=%s", url)
}

func (c *RedditCounter) Count(r *Request) (count int) {
	var resp *http.Response
	var dec *json.Decoder
	var err error
	var data *redditFeed
	var url = r.Url()

	if resp, err = http.Get(c.urlFor(url)); err != nil {
		return 0
	}

	dec = json.NewDecoder(resp.Body)
	if err = dec.Decode(&data); err != nil {
		return 0
	}

	if len(data.Data.Children) > 0 {
		count = data.Data.Children[0].Data.Score
	}
	return
}

func init() {
	counters["reddit"] = new(RedditCounter)
}
