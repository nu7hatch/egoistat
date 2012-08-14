package egoistat

import (
	"encoding/json"
	"net/http"
	"strings"
)

type googlePlusFeed struct {
	Result struct {
		Metadata struct {
			GlobalCounts struct {
				Count float64
			}
		}
	}
}

type googlePlusRequest struct {
	Method     string                   `json:"method"`
	Id         string                   `json:"id"`
	Key        string                   `json:"key"`
	RPCVersion string                   `json:"jsonrpc"`
	APIVersion string                   `json:"apiVersion"`
	Params     *googlePlusRequestParams `json:"params"`
}

type googlePlusRequestParams struct {
	Nolog   bool   `json:"nolog"`
	Id      string `json:"id"`
	Source  string `json:"source"`
	UserId  string `json:"userId"`
	GroupId string `json:"groupId"`
}

type GooglePlusCounter struct{}

func (c *GooglePlusCounter) url() string {
	return "https://clients6.google.com/rpc?key=AIzaSyCKSbrvQasunBoV16zDH9R33D88CeLr9gQ"
}

func (c *GooglePlusCounter) postDataFor(url string) string {
	b, err := json.Marshal([]*googlePlusRequest{
		&googlePlusRequest{
			"pos.plusones.get", "p", "p", "2.0", "v1",
			&googlePlusRequestParams{true, url, "widget", "@viewer", "@self"},
		},
	})
	if err != nil {
		return ""
	}
	return string(b)
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
