package setup

import (
	"io/fs"
	"log"
	"path/filepath"
	"strings"

	"github.com/FelixSnz/mail-indexer/internal/email"
	"github.com/FelixSnz/mail-indexer/pkg/zinc"
)

func UploadDocuments(emails_path string) {

	var root string

	if strings.Contains(emails_path, "\\") {
		splitted_path := strings.Split(emails_path, "\\")
		root = splitted_path[len(splitted_path)-1]
	} else {
		root = emails_path
	}

	filepath.WalkDir(emails_path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatalf(err.Error())
		}
		if !d.IsDir() {

			email := email.New(path, root)
			json := email.GetJson()
			zinc.AddDoc(json)
		}
		return nil
	})

}
