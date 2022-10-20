package zinc

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func Search(field, pattern string) string {

	index_url := "http://localhost:4080/api/emails/"
	api_call := "_search"
	fmt.Println("HTTP JSON POST URL:", index_url+api_call)

	query := fmt.Sprintf(`{
        "search_type": "querystring",
        "query":
        {
            "term": "%s:%s"
        },
        "from": 0,
        "max_results": 20,
        "_source": []
    }`, field, pattern)
	req, err := http.NewRequest("POST", index_url+api_call, strings.NewReader(query))
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
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	log.Println(resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}
