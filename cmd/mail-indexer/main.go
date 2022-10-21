package main

import (
	"github.com/FelixSnz/mail-indexer/pkg/zinc"
)

func main() {

	//test to update single document
	asd := `{
		"date": "2002-02-05T09:43:01Z",
		"bcc": [],
		"cc": [],
		"content": "Conference call with UBS",
		"encoding": "7bit",
		"from": "john.zufferli@enron.com",
		"id": "<28390332.1075842024763.JavaMail.evans@thyme>",
		"origin": "zufferli-j",
		"subfolder": "calendar",
		"subject": "",
		"to": [],
		"type": "text/plain; charset=us-ascii",
		"mimeversion": "1.0"
	  }`

	zinc.AddDoc(asd)
	//fmt.Print("este fue el response: ")
	//fmt.Println(zinc.Search("content", "manipulated"))

	//fmt.Println("este fue el str date: " + email.CleanDate("Wed, 6 Dec 2000 22:23:00 -0800 (PST)").String())

	//setup.Run(`api\zincsearch\email_samples`, "testmails.json")

	//setup.Run(`api\zincsearch\email_samples`, "test-date.json")
}
