package main

import (
	"calcium/pkg/email"
	"calcium/pkg/zinc"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	args := os.Args[1:]

	if len(args) == 2 {

		index_name := args[0]
		mails_path := args[1]

		var root string

		if strings.Contains(mails_path, "\\") {
			splitted_path := strings.Split(mails_path, "\\")
			root = splitted_path[len(splitted_path)-1]
		} else {
			root = mails_path
		}

		zinc_index := zinc.NewIndex(index_name)

		filepath.WalkDir(mails_path, func(path string, de fs.DirEntry, err error) error {
			if err != nil {
				fmt.Printf("error: couldn't iterate over files at path %s\n", mails_path)
				log.Fatal(err)
			}
			if !de.IsDir() {

				fmt.Printf("path: %s\n", path)

				email_obj := email.New(path, root)
				json := email_obj.GetJson()
				zinc_index.Post("_doc", json)
			}
			return nil
		})

	} else {
		fmt.Printf("error: invalid args length, expected: 2, recv: %d\n", len(args))
		os.Exit(1)
	}

}
