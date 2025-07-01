package main

import (
	"encoding/json"
	"fmt"
)

// Task: parse raw JSON response
const rawResp = `
{
    "header": {
        "code": 0,
        "message": ""
    },
    "data": [{
        "type": "user",
        "id": 123,
        "attributes": {
            "email": "bob@google.com",
            "article_ids": [5, 8, 12]
        }
    }]
}
`

type Response struct {
	Header struct {
		Code    int    `json:"code"`
		Message string `json:"message,omitempty"`
	} `json:"header"`
	Data []struct {
		Type       string `json:"type"`
		ID         int    `json:"id"`
		Attributes struct {
			Email    string `json:"email"`
			Articles []int  `json:"article_ids"`
		} `json:"attributes"`
	} `json:"data,omitempty"`
}

func main() {
	var r Response

	if err := json.Unmarshal([]byte(rawResp), &r); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(r)
}
