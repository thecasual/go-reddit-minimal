package reddit

import (
	"fmt"
	"testing"
)

// Build Client
func TestBuildClient(t *testing.T) {
	client := NewClient(userName, password, redditUrl, clientID, clientSecret)
	if client.token == "" {
		t.Errorf("Token is empty")
	} else {
		fmt.Println("Building Test Client: Success")
	}
}

// Get Self
func TestGetMe(t *testing.T) {
	client := NewClient(userName, password, redditUrl, clientID, clientSecret)
	got := client.GetMe()
	if got == "" {
		t.Errorf("Error /api/v1/me")
	} else {
		fmt.Println("Query /api/v1/me: Success")
	}

	fmt.Printf("Evening, %s", client.GetMe())
}

// Search
func TestSubredditsearch(t *testing.T) {
	client := NewClient(userName, password, redditUrl, clientID, clientSecret)
	client.Search()
}
