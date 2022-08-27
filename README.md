#

### Get:
```go get -u github.com/thecasual/go-reddit-minimal```
#


### Example:

```
package main

import (
	"fmt"
    "github.com/thecasual/go-reddit-minimal/reddit"

)

const (
	url          = ""
	clientID     = ""
	clientSecret = ""
	userName     = ""
	password     = ""
)

func main() {
	client := NewClient(userName, password, url, clientID, clientSecret)
	fmt.Printf("Evening, %s", client.GetMe())
}

```
# 
