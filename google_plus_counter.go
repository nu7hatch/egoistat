package egoist

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type GooglePlusCounter struct{}

type googlePlusFeed struct {
	Result googlePlusFeedResult
}

type googlePlusFeedResult struct {
	Metadata googlePlusFeedResultMetadata
}

type googlePlusFeedResultMetadata struct {
	GlobalCounts googlePlusFeedResultMetadataCounts
}

type googlePlusFeedResultMetadataCounts struct {
	Count float64
}

func (c *GooglePlusCounter) url() string {
	return "https://clients6.google.com/rpc?key=AIzaSyCKSbrvQasunBoV16zDH9R33D88CeLr9gQ"
}

func (c *GooglePlusCounter) postDataFor(url string) string {
	return fmt.Sprintf(
		"[{\"method\":\"pos.plusones.get\",\"id\":\"p\",\"params\":{"+
			"\"nolog\":true,\"id\":\"%s\",\"source\":\"widget\",\"userId\":"+
			"\"@viewer\",\"groupId\":\"@self\"},\"jsonrpc\":\"2.0\",\"key\":"+
			"\"p\",\"apiVersion\":\"v1\"}]", url,
	)
}

func (c *GooglePlusCounter) Count(r *Request) int {
	var resp *http.Response
	var dec *json.Decoder
	var err error
	var data []googlePlusFeed
	var postData = strings.NewReader(c.postDataFor(r.Url()))

	if resp, err = http.Post(c.url(), "application/json", postData); err != nil {
		return 0
	}

	dec = json.NewDecoder(resp.Body)
	if err = dec.Decode(&data); err != nil {
		return 0
	}
	if len(data) == 0 {
		return 0
	}

	return int(data[0].Result.Metadata.GlobalCounts.Count)
}

func init() {
	counters["plusone"] = new(GooglePlusCounter)
}
