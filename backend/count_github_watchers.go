package egoistat

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const githubFeedUrl = "https://api.github.com/repos/%s"

type githubFeed struct {
	Watchers int
}

func CountGithubWatchers(r *Request) *Result {
	var resp *http.Response
	var dec *json.Decoder
	var err error
	var feed githubFeed
	var repoName = r.Option("github_repo")
	var feedUrl = fmt.Sprintf(githubFeedUrl, repoName)

	if resp, err = http.Get(feedUrl); err != nil {
		return Empty
	}

	dec = json.NewDecoder(resp.Body)
	if err = dec.Decode(&feed); err != nil {
		return Empty
	}

	return &Result{Points: feed.Watchers}
}

func init() {
	RegisterCounter("github", CountGithubWatchers)
}
