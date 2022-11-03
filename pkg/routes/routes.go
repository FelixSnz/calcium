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
	Keyword string `json:"keyword"`
	From    string `json:"from"`
	To      string `json:"to"`
}

func SearchEmails(w http.ResponseWriter, r *http.Request) {

	var searchReq SearchReq

	reqBody, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "insert a valid ")
	}

	zinc_index := zinc.NewIndex("enron_mails")

	json.Unmarshal(reqBody, &searchReq)

	fmt.Printf("keyword: %s\n", searchReq.Keyword)
	fmt.Printf("from: %s\n", searchReq.From)
	fmt.Printf("to: %s\n", searchReq.To)

	searchQuery := fmt.Sprintf(`{
				"search_type": "querystring",
				"query": {
					"term": "+content:\"%s\" +date:>\"%s\" +date:<\"%s\""
				},
		
					"_source":[]
			}`, searchReq.Keyword, searchReq.From, searchReq.To)

	zinc_resp := zinc_index.Search(searchQuery)

	w.Header().Set("Content-Type", "text/plain")
	w.Write(zinc.GetOnlySource[email.JsonStruct](zinc_resp))
	//w.Write((zinc_resp))
}
