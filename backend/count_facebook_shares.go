package egoistat

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const facebookFeedUrl = "http://graph.facebook.com/?ids=%s"

type facebookFeed struct {
	Shares int
}

func CountFacebookShares(r *Request) *Result {
	var (
		resp *http.Response
		dec *json.Decoder
		err error
		feed map[string]facebookFeed
		url = r.Url()
		feedUrl = fmt.Sprintf(facebookFeedUrl, url)
	)
	
	if resp, err = http.Get(feedUrl); err != nil {
		return Empty
	}

	dec = json.NewDecoder(resp.Body)
	if err = dec.Decode(&feed); err != nil {
		return Empty
	}

	feedForUrl, ok := feed[url]
	if !ok {
		return Empty
	}

	return &Result{Points: feedForUrl.Shares}
}

func init() {
	RegisterCounter("facebook", CountFacebookShares)
}
