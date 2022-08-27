package reddit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Define client
type redditClient struct {
	c                *http.Client
	userName         string
	password         string
	url              string
	clientID         string
	clientSecret     string
	token            string
	authorizationUrl string
	oauthUrl         string
}

// Build client
func NewClient(userName string,
	password string,
	url string,
	clientID string,
	clientSecret string) *redditClient {

	c := &http.Client{}

	client := &redditClient{
		c:                c,
		userName:         userName,
		password:         password,
		url:              url,
		clientID:         clientID,
		clientSecret:     clientSecret,
		token:            "",
		authorizationUrl: "https://www.reddit.com/api/v1/authorize",
		oauthUrl:         "https://oauth.reddit.com",
	}

	client.token = GetToken(*client)
	return client
}

// HTTP Request
func doRequest(client redditClient, method string, url string, body string) (*http.Response, error) {
	bodyByte := []byte(body)

	req, _ := http.NewRequest(method, url, bytes.NewBuffer(bodyByte))

	// Token exists
	if client.token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("bearer %s", client.token))

	} else {
		// Getting token via oauth flow
		req.SetBasicAuth(client.clientID, client.clientSecret)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.c.Do(req)
	codeHandled := []int{429, 200}

	for {

		if err == nil {

			// Pause for 429s
			if resp.StatusCode == 429 {

				time.Sleep(5 * time.Second)
				resp, err = client.c.Do(req)
				continue
			}

			if resp.StatusCode == 200 {
				break
			}

			if !contains(codeHandled, resp.StatusCode) {
				break
			}

		}
	}

	return resp, err
}

// Return bearer token
func GetToken(client redditClient) string {
	var stringdata = fmt.Sprintf("grant_type=password&username=%s&password=%v", client.userName, client.password)
	resp, _ := doRequest(client, "POST", fmt.Sprintf("%s/api/v1/access_token", client.url), stringdata)
	json := processJSON(resp)
	return json["access_token"].(string)
}

// Process JSON in http.Response
func processJSON(resp *http.Response) map[string]interface{} {
	body, _ := io.ReadAll(resp.Body)
	body_str := string(body)
	var objmap map[string]interface{}
	if err := json.Unmarshal([]byte(body_str), &objmap); err != nil {
		log.Fatal(err)
	}
	return objmap
}

// api/v1/me
func GetMe(client redditClient) string {
	req, _ := doRequest(client, "GET", fmt.Sprintf("%s/api/v1/me", client.oauthUrl), "")
	json := processJSON(req)
	return json["name"].(string)
}
