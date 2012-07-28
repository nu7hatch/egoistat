package egoist

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GithubCounter struct{}

type githubFeed struct {
	Watchers int
}

func (c *GithubCounter) urlFor(repoName string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s", repoName)
}

func (c *GithubCounter) Count(r *Request) int {
	var resp *http.Response
	var dec *json.Decoder
	var err error
	var data githubFeed
	var repoName = r.Option("github_repo")

	if resp, err = http.Get(c.urlFor(repoName)); err != nil {
		return 0
	}

	dec = json.NewDecoder(resp.Body)
	if err = dec.Decode(&data); err != nil {
		return 0
	}

	return data.Watchers
}

func init() {
	counters["github"] = new(GithubCounter)
}
