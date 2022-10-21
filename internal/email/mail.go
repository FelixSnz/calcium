// Package for describing email object thus handling email files
// with the general porpuse of generating a the json format of the email file
package email

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/FelixSnz/mail-indexer/internal/utils"
)

// defines Email file as an object to store its information as json format
type JsonStruct struct {
	Origin      string    `json:"origin"`
	SubFolder   string    `json:"subfolder"`
	Id          string    `json:"id"`
	Date        time.Time `json:"date"`
	From        string    `json:"from"`
	To          []string  `json:"to"`
	Cc          []string  `json:"cc"`
	Bcc         []string  `json:"bcc"`
	Subject     string    `json:"subject"`
	Version     string    `json:"mimeversion"`
	ContentType string    `json:"type"`
	Encoding    string    `json:"encoding"`
	Content     string    `json:"content"`
}

type Email struct {
	path   string
	root   string
	states [10]string // defines scan states when reading email file line by line
}

// constructor method to create a new email object
func New(path, root string) *Email {
	email := new(Email)
	email.path = path
	email.root = root

	email.states = [10]string{
		"Message-ID: ",
		"Date: ",
		"From: ",
		"To: ",
		"Cc: ",
		"Bcc: ",
		"Subject: ",
		"Mime-Version: ",
		"Content-Type: ",
		"Content-Transfer-Encoding: ",
	}
	return email
}

// returns the message of the email
func (email Email) GetMessage() string {

	raw_data, err := os.ReadFile(email.path)

	if err != nil {
		fmt.Printf("error: Couldn't read file at path: %s\n", email.path)
		log.Fatal(err)
	}

	//to get the message content of the email file, separates
	//the whole content by a regular expression
	reg_delimeter := regexp.MustCompile(`X-FileName\s*(.*?)\s*\n`)
	content := reg_delimeter.Split(string(raw_data), -1)[1]
	return content

}

// by a given email file line, returns the scan state if found, otherwise, returns the provided previus state
func (email Email) GetState(line, prev_state string) string {

	for i := 0; i < len(email.states); i++ {

		if strings.HasPrefix(line, email.states[i]) {
			return email.states[i]
		}
	}

	return prev_state

}

// by the path of the emil instance, returns a json string
func (email Email) GetJson() string {
	readFile, err := os.Open(email.path)
	if err != nil {
		fmt.Printf("couldn't open email file at path: %s\n", email.path)
		panic(err)

	}

	var ID string
	var date string
	var from_mail string
	var to_mails string
	var cc_mails string
	var subject string
	var mime_version string
	var content_type string
	var encoding string
	var bcc_mails string

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	prev_state := ""

	for fileScanner.Scan() {

		line := fileScanner.Text()

		state := email.GetState(line, prev_state)

		if !strings.HasPrefix(line, "X-") {

			switch state {

			case "Message-ID: ":
				ID += line

			case "Date: ":
				date += line

			case "From: ":
				from_mail += line

			case "To: ":
				to_mails += line

			case "Cc: ":
				cc_mails += line

			case "Subject: ":
				subject += line

			case "Mime-Version: ":
				mime_version += line

			case "Content-Type: ":
				content_type += line

			case "Content-Transfer-Encoding: ":
				encoding += line

			case "Bcc: ":
				bcc_mails += line

			default:

			}

		} else {
			break
		}

		prev_state = state

	}

	readFile.Close()

	splited_path := strings.Split(email.path, "\\")

	root_idx := utils.IndexOf(email.root, splited_path)

	if root_idx == -1 {
		log.Fatal("couldn't extract the root index")
	}

	sub_folder := strings.Join(splited_path[root_idx+2:len(splited_path)-1], `\`)
	email_msg := email.GetMessage()

	json_str, err := json.Marshal(JsonStruct{
		Origin:      splited_path[root_idx+1],
		SubFolder:   sub_folder,
		Id:          strings.TrimPrefix(ID, "Message-ID: "),
		Date:        CleanDate(date),
		From:        strings.TrimPrefix(from_mail, "From: "),
		To:          GetMails(strings.TrimPrefix(to_mails, "To: ")),
		Cc:          GetMails(strings.TrimPrefix(cc_mails, "Cc: ")),
		Bcc:         GetMails(strings.TrimPrefix(bcc_mails, "Bcc: ")),
		Subject:     strings.TrimPrefix(subject, "Subject: "),
		Version:     strings.TrimPrefix(mime_version, "Mime-Version: "),
		ContentType: strings.TrimPrefix(content_type, "Content-Type: "),
		Encoding:    strings.TrimPrefix(encoding, "Content-Transfer-Encoding: "),
		Content:     email_msg,
	})

	if err != nil {
		fmt.Printf("error: Couldn't create json file at path: %s\n", email.path)
		log.Fatal(err)
	}

	return string(json_str)

}

// by a given string of email adresses, removes spaces, tabs, new lines and returns an array
// where each element is an email address
func GetMails(mails_string string) []string {

	return strings.FieldsFunc(
		strings.TrimFunc(
			strings.Replace(
				strings.Replace(mails_string,
					" ", "", -1), "\t", "", -1), func(c rune) bool {
				return c == 9 || c == 10 || c == 11
			}),
		func(c rune) bool {
			return c == ','
		})

}

// by a raw date extracted from the email content, cleans the format and returns
// the date in the default time type for go
func CleanDate(raw_date string) time.Time {

	str_date := strings.TrimPrefix(raw_date, "Date: ")
	day := str_date[5:7]
	no_zero_pattern := `\d{1} `
	match, err := regexp.MatchString(no_zero_pattern, day)

	if err != nil {
		fmt.Println("error matching day string")
		log.Fatal(err)
	}

	if match {

		index := 5
		str_date = str_date[:index] + "0" + str_date[index:]
	}

	num_zone_regex := regexp.MustCompile(`-\s*(\d{4}) |\(|\)`)
	str_date = num_zone_regex.ReplaceAllString(str_date, "")
	date, err := time.Parse(time.RFC1123, str_date)

	if err != nil {
		fmt.Println("error parsing date")
		log.Fatal(err)
	}
	return date
}
