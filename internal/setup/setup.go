package setup

import (
	"fmt"

	"github.com/FelixSnz/mail-indexer/internal/email"
)

func Run(emails_path, json_savename string) {
	fmt.Println("Running setup data")
	email.DirToJson(emails_path, json_savename)
}
