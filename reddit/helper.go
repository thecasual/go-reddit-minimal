package reddit

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Check if contains
func contains(s []int, str int) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// Process JSON in http.Response
func processJSONReq(resp *http.Response) map[string]interface{} {
	body, _ := io.ReadAll(resp.Body)
	body_str := string(body)
	var objmap map[string]interface{}
	if err := json.Unmarshal([]byte(body_str), &objmap); err != nil {
		log.Fatal(err)
	}
	return objmap
}
