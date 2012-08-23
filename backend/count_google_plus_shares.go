package egoistat

import (
	"encoding/json"
	"net/http"
	"strings"
)

const googlePlusFeedUrl = "https://clients6.google.com/rpc?key=AIzaSyCKSbrvQasunBoV16zDH9R33D88CeLr9gQ"

type googlePlusFeed struct {
	Result struct {
		Metadata struct {
			GlobalCounts struct {
				Count float64
			}
		}
	}
}

type googlePlusRequestData struct {
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

func newGooglePlusRequestData(url string) string {
	b, err := json.Marshal([]*googlePlusRequestData{
		&googlePlusRequestData{
			"pos.plusones.get", "p", "p", "2.0", "v1",
			&googlePlusRequestParams{true, url, "widget", "@viewer", "@self"},
		},
	})
	if err != nil {
		return ""
	}
	return string(b)
}

func CountGooglePlusShares(r *Request) *Result {
	var (
		resp     *http.Response
		dec      *json.Decoder
		err      error
		feed     []googlePlusFeed
		postData = strings.NewReader(newGooglePlusRequestData(r.Url()))
	)

	if resp, err = http.Post(googlePlusFeedUrl, "application/json", postData); err != nil {
		return Empty
	}

	dec = json.NewDecoder(resp.Body)
	if err = dec.Decode(&feed); err != nil {
		return Empty
	}
	if len(feed) == 0 {
		return Empty
	}

	return &Result{Points: int(feed[0].Result.Metadata.GlobalCounts.Count)}
}

func init() {
	RegisterCounter("plusone", CountGooglePlusShares)
}
