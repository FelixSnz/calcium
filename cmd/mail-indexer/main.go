package main

import (
	"fmt"

	"github.com/FelixSnz/mail-indexer/pkg/zinc"
)

func main() {
	fmt.Print("este fue el response: ")
	fmt.Println(zinc.Search("content", "manipulated"))
}
