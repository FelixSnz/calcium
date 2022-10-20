package main

import (
	"github.com/FelixSnz/mail-indexer/internal/setup"
)

func main() {
	//fmt.Print("este fue el response: ")
	//fmt.Println(zinc.Search("content", "manipulated"))

	//fmt.Println("este fue el str date: " + email.CleanDate("Wed, 6 Dec 2000 22:23:00 -0800 (PST)").String())

	setup.Run(`api\zincsearch\email_samples`, "testmails.json")

	//setup.Run(`api\zincsearch\email_samples`, "test-date.json")
}
