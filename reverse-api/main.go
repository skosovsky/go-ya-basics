package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Response struct {
	Header struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"header"`
	Data []struct {
		Type       string `json:"type"`
		ID         int    `json:"id"`
		Attributes struct {
			Email      string `json:"email"`
			ArticleIDs []int  `json:"articleIDs"`
		} `json:"attributes"`
	} `json:"data"`
}

func ReadResponse(rawResp string) (Response, error) {
	var resp Response

	err := json.Unmarshal([]byte(rawResp), &resp)
	if err != nil {
		err = fmt.Errorf("unmarshal error %w", err)
		log.Println(err)

		return Response{}, err
	}

	return resp, nil
}

func main() {
	rawResp := `{
    "header": {
        "code": 0,
        "message": ""
    },
    "data": [{
        "type": "user",
        "id": 100,
        "attributes": {
            "email": "bob@yandex.ru",
            "article_ids": [10, 11, 12]
        }
    }]
}`

	response, err := ReadResponse(rawResp)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(response) //nolint:forbidigo // example
}
