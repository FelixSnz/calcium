package main

import (
	"github.com/FelixSnz/mail-indexer/internal/setup"
)

func main() {
	//fmt.Println(zinc.Search("Content", "manipulated"))

	setup.Run(`assets\email_samples`, "email_samples.json")
}
