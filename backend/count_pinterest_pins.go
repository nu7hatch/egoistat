package egoistat

import (
	"encoding/json"
	"fmt"
	"strings"
	"io/ioutil"
	"net/http"
)

const pinterestFeedUrl = "http://api.pinterest.com/v1/urls/count.json?url=%s"

type pinterestFeed struct {
	Count int
}

func CountPinterestPins(r *Request) *Result {
	var (
		resp *http.Response
		err error
		feed pinterestFeed
		bodyBytes []byte
		body string
		feedUrl = fmt.Sprintf(pinterestFeedUrl, r.Url())
	)

	if resp, err = http.Get(feedUrl); err != nil {
		return Empty
	}

	bodyBytes, err = ioutil.ReadAll(resp.Body)
	body = string(bodyBytes)
	body = strings.Replace(body, "receiveCount(", "", -1)
	body = strings.Replace(body, ")", "", -1)

	if err = json.Unmarshal([]byte(body), &feed); err != nil {
		return Empty
	}

	return &Result{Points: feed.Count}
}

func init() {
	RegisterCounter("pinterest", CountPinterestPins)
}
