// package to call zincsearch api methods
package zinc

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type ErrorResponse struct {
	Err string `json:"error"`
}

type zincIndex struct {
	name string
}

func NewIndex(name string) *zincIndex {
	index := new(zincIndex)
	index.name = name

	return index
}

// Calls post methods to upload documents or send search queries
//
// Query examples:
//
// search by date query:
//
//	`{
//		"search_type": "querystring",
//		"query": {
//			"term": "+date:>\"2000-01-09T15:01:00-07:00\" +date:<\"2000-12-09T15:01:00-07:00\""
//		},
//
//			"_source":["date"]
//	}`
//
// search by word query:
//
//	{
//		"search_type": "querystring",
//		"query":
//		{
//			"term": "content:manipulated"
//		},
//		"from": 0,
//		"max_results": 20,
//		"_source": []
//	}
func (index zincIndex) Post(method, json_str string) []byte {

	index_url := fmt.Sprintf(`http://localhost:4080/api/%s/`, index.name)

	req, err := http.NewRequest("POST", index_url+method, strings.NewReader(json_str))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	//log.Printf("resp: %d\n", resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {

		var err_resp ErrorResponse
		json.Unmarshal(body, &err_resp)

		fmt.Println("error from zincsearch: ")
		log.Fatal(err_resp.Err)

	}

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)

	// fmt.Println(body)
	return body

}
