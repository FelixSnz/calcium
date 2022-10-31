package routes

import (
	"calcium/pkg/zinc"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SearchReq struct {
	Keyword string `json:"Keyword"`
	From    string `json:"From"`
	To      string `json:"To"`
}

func SearchEmails(w http.ResponseWriter, r *http.Request) {

	var searchReq SearchReq

	reqBody, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "insert a valid ")
	}

	zinc_index := zinc.NewIndex("enron_mails")

	json.Unmarshal(reqBody, &searchReq)

	searchQuery := fmt.Sprintf(`{
				"search_type": "querystring",
				"query": {
					"term": "+Content:\"%s\" +date:>\"%s\" +date:<\"%s\""
				},
		
					"_source":["date"]
			}`, searchReq.Keyword, searchReq.From, searchReq.To)

	zinc_resp := zinc_index.Post("_search", searchQuery)

	w.Header().Set("Content-Type", "text/plain")
	w.Write(zinc_resp)
}
