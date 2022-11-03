// package to call zincsearch api methods
package zinc

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type FullResponse[sourceType interface{}] struct {
	Took     int              `json:"took"`
	TimedOut bool             `json:"timed_out"`
	Shards   Shard            `json:"_shards"`
	Hits     Hits[sourceType] `json:"hits"`
}

type Shard struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Skipped    int `json:"skipped"`
	Failed     int `json:"failed"`
}

type HitTotalValue struct {
	Value int `json:"value"`
}

type Hits[sourceType interface{}] struct {
	Total    HitTotalValue          `json:"total"`
	MaxScore float64                `json:"max_score"`
	Hits     []IndexDoc[sourceType] `json:"hits"`
}

type IndexDoc[sourceType interface{}] struct {
	Index     string     `json:"_index"`
	Type      string     `json:"_type"`
	Id        string     `json:"_id"`
	Score     float64    `json:"_score"`
	Timestamp time.Time  `json:"@timestamp"`
	Source    sourceType `json:"_source"`
}

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
func (index zincIndex) Search(query string) []byte {

	index_url := fmt.Sprintf(`http://localhost:4080/api/%s/`, index.name)

	req, err := http.NewRequest("POST", index_url+"_search", strings.NewReader(query))
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

	return body

}

func GetOnlySource[sourceType interface{}](ZincResponse []byte) []byte {

	var fullResp FullResponse[sourceType]
	err := json.Unmarshal(ZincResponse, &fullResp)

	if err != nil {
		fmt.Println("error unmarshaling the body from zinc response")
		log.Fatal(err)
	}
	var source_data []sourceType

	for _, elm := range fullResp.Hits.Hits {
		source_data = append(source_data, elm.Source)
	}

	source, err := json.Marshal(source_data)

	if err != nil {
		fmt.Println("error marshaling emails array into json")
		log.Fatal(err)
	}

	return source

}
