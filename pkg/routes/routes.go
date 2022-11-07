package routes

import (
	"calcium/pkg/email"
	"calcium/pkg/zinc"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SearchReq struct {
	Type     string   `json:"type"`
	Keywords []string `json:"keywords"`
	From     string   `json:"from"`
	To       string   `json:"to"`
}

func SearchEmails(w http.ResponseWriter, r *http.Request) {

	var searchReq SearchReq

	reqBody, err := io.ReadAll(r.Body)

	fmt.Printf("reqbody: %s\n", string(reqBody))

	if err != nil {
		fmt.Fprintf(w, "insert a valid ")
	}

	zinc_index := zinc.NewIndex("enron_emails")

	json.Unmarshal(reqBody, &searchReq)

	fmt.Printf("to: %s\n", searchReq.Type)
	fmt.Printf("keyword: %s\n", searchReq.Keywords)
	fmt.Printf("from: %s\n", searchReq.From)
	fmt.Printf("to: %s\n", searchReq.To)

	var from_date_term string = ""
	var to_date_term string = ""

	if searchReq.From != "" {
		from_date_term = fmt.Sprintf(` +date:>\"%s\"`, searchReq.From)

	}
	if searchReq.To != "" {
		to_date_term = fmt.Sprintf(` +date:<\"%s\"`, searchReq.To)

	}

	for i, keyword := range searchReq.Keywords {
		searchReq.Keywords[i] = "+" + searchReq.Type + ":" + keyword
	}

	keywords_str := fmt.Sprintf("%v", searchReq.Keywords)

	searchQuery := fmt.Sprintf(`{
				"search_type": "querystring",
				"query": {
					"term": "%s%s%s"
				},
				"from": 0,
				"max_results": 600000,
				"_source":[]
			}`, keywords_str[1:len(keywords_str)-1], from_date_term, to_date_term)

	fmt.Printf("query: %s\n", searchQuery)
	zinc_resp := zinc_index.Post("_search", searchQuery)

	fmt.Printf("respbody: %s\n", string(zinc_resp))

	w.Header().Set("Content-Type", "text/plain")
	w.Write(zinc.GetOnlySource[email.JsonStruct](zinc_resp))
	//w.Write((zinc_resp))
}
