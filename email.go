package main

import (
	"fmt"
	"log"
	"net/smtp"
	"encoding/json"
	"io/ioutil"
)

func email(string message){

	data, err := ioutil.ReadFile("mail_config.json")
	if err != nil {
		fmt.Print(err)
	}

	type MailConfig struct {
                recipients []string `json:"recipients"`
                sender string `json:"sender"`
                hostname []byte `json:"hostname"`
                password string `json:"password"`
		port string `json:"port"`
        }

	var obj MailConfig
	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println ("error:", err)
	}

	hostname := "smtp.gmail.com"
	auth := smtp.PlainAuth("", obj.sender, obj.password, hostname)
	err = smtp.SendMail(obj.hostname+":"+obj.port, auth, obj.sender, obj.recipients,[]byte(message))
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error: %s",err)
	}

}

func main(){
	email("My first Golang Email")
}
