// Package for describing email object thus handling email files
// with the general porpuse of generating a the json format of the email file
package email

import (
	"bufio"
	"calcium/pkg/utils"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

// defines Email file relevant content as an object to store its information as json format
type JsonStruct struct {
	Origin    string    `json:"origin"`
	SubFolder string    `json:"subfolder"`
	Date      time.Time `json:"date"`
	Subject   string    `json:"subject"`
	From      string    `json:"from"`
	To        []string  `json:"to"`
	Cc        []string  `json:"cc"`
	Content   string    `json:"content"`

	// Bcc         []string  `json:"bcc"`
	// Id          string    `json:"id"`
	// Version     string    `json:"mimeversion"`
	// ContentType string    `json:"type"`
	// Encoding    string    `json:"encoding"`

}

// struct to handle email files
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
func (email Email) GetContent() string {

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

	var date string
	var from_mail string
	var to_mails string
	var cc_mails string
	var subject string

	//var ID string
	// var content_type string
	// var encoding string
	// var bcc_mails string

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	prev_state := ""

	for fileScanner.Scan() {

		line := fileScanner.Text()

		state := email.GetState(line, prev_state)

		if !strings.HasPrefix(line, "X-") {

			switch state {

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

			// case "Message-ID: ":
			// 	ID += line

			// case "Mime-Version: ":
			// 	mime_version += line

			// case "Content-Type: ":
			// 	content_type += line

			// case "Content-Transfer-Encoding: ":
			// 	encoding += line

			// case "Bcc: ":
			// 	bcc_mails += line

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

	//fmt.Printf("splitted path: %v", splited_path)

	sub_folder := strings.Join(splited_path[root_idx+2:len(splited_path)-1], `\`)
	content := email.GetContent()

	json_str, err := json.Marshal(JsonStruct{
		Origin:    splited_path[root_idx+1],
		SubFolder: sub_folder,
		Date:      CleanDate(date),
		Subject:   strings.TrimPrefix(subject, "Subject: "),
		From:      strings.TrimPrefix(from_mail, "From: "),
		To:        GetMails(strings.TrimPrefix(to_mails, "To: ")),
		Cc:        GetMails(strings.TrimPrefix(cc_mails, "Cc: ")),
		Content:   content,

		// Id:          strings.TrimPrefix(ID, "Message-ID: "),
		// Bcc:         GetMails(strings.TrimPrefix(bcc_mails, "Bcc: ")),
		// Version:     strings.TrimPrefix(mime_version, "Mime-Version: "),
		// ContentType: strings.TrimPrefix(content_type, "Content-Type: "),
		// Encoding:    strings.TrimPrefix(encoding, "Content-Transfer-Encoding: "),
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

	str_date = strings.Split(str_date, " (")[0]

	date, err := time.Parse(time.RFC1123Z, str_date)

	if err != nil {
		fmt.Println("error parsing date")
		log.Fatal(err)
	}
	return date
}
