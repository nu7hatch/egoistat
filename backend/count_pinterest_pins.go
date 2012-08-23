package egoistat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"bytes"
)

const pinterestFeedUrl = "http://api.pinterest.com/v1/urls/count.json?url=%s"

type pinterestFeed struct {
	Count int
}

func CountPinterestPins(r *Request) *Result {
	var (
		resp    *http.Response
		err     error
		feed    pinterestFeed
		body    []byte
		feedUrl = fmt.Sprintf(pinterestFeedUrl, r.Url())
	)

	if resp, err = http.Get(feedUrl); err != nil {
		return Empty
	}
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return Empty
	}

	body = bytes.Replace(body, []byte("receiveCount("), nil, 1)
	body = bytes.Replace(body, []byte(")"), nil, 1)

	if err = json.Unmarshal([]byte(body), &feed); err != nil {
		return Empty
	}

	return &Result{Points: feed.Count}
}

func init() {
	RegisterCounter("pinterest", CountPinterestPins)
}
