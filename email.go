package main

import (
	"fmt"
	"log"
	"net/smtp"
	"encoding/json"
	"io/ioutil"
)

func main(){

	data, err := ioutil.ReadFile("mail_config.json")
	if err != nil {
		fmt.Print(err)
	}

	type MailConfig struct {
                recipients []string `json:"recipients"`
                sender string `json:"sender"`
                message []byte `json:"message"`
                password string `json:"password"`
        }

	var obj MailConfig
	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println ("error:", err)
	}

	hostname := "smtp.gmail.com"
	auth := smtp.PlainAuth("", obj.sender, obj.password, hostname)
	err = smtp.SendMail("smtp.gmail.com:25", auth, obj.sender, obj.recipients, obj.message)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error: %s",err)
	}

}
