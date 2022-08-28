package reddit

import (
	"encoding/json"
	"fmt"
	"io"
)

// api/v1/me
func (client redditClient) GetMe() string {
	req, _ := doRequest(client, "GET", fmt.Sprintf("%s/api/v1/me", client.oauthUrl), "")
	json := processJSONReq(req)
	return json["name"].(string)
}

// Return []interface{} of search entries
func (client redditClient) Search() []interface{} {

	// Do request
	url := fmt.Sprintf("%s/search", client.oauthUrl)
	url = fmt.Sprintf("%s?q=subreddit:netsec&limit=100", url)
	req, _ := doRequest(client, "GET", url, "")

	// Process response
	body, _ := io.ReadAll(req.Body)
	json_all := &RedditSearchResponse{}
	json.Unmarshal(body, json_all)

	// Build return []interface{}
	var allReq []interface{}
	for _, entry := range json_all.Data.Children {
		allReq = append(allReq, entry.Data)
	}

	return allReq
}
