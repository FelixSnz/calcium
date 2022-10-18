package setup

import (
	"fmt"

	"github.com/FelixSnz/mail-indexer/internal/email"
)

func Run(emails_path string) {
	fmt.Println("Running setup data")
	email.DirToJson(emails_path)

}
