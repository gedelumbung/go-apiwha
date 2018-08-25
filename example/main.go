package main

import (
	"fmt"

	apiwha "github.com/gedelumbung/go-apiwha"
)

const (
	url = "http://panel.apiwha.com"
	key = "YOUR KEY HERE"
)

func main() {
	wa := apiwha.Init(url, key)
	credit := wa.Credit()
	fmt.Println(credit.Result)

	messageParams := &apiwha.ApiWhaMessagesParams{
		Type: "OUT",
	}
	messages := wa.Messages(messageParams)
	fmt.Printf("%#v", messages.Result)

	sendMessageParams := &apiwha.ApiWhaSendMessageParams{
		Number: "DESTINATION PHONE NUMBER",
		Text:   "Send from Golang Library after check credit & pull messages",
	}
	sendMessage := wa.SendMessage(sendMessageParams)
	fmt.Println(sendMessage.Result)
}
