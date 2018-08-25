

# Golang APIWHA Library (Unofficial)
So much thanks to APIWHA, because already made an awesome Whatsapp API and also give $10 free credit :D. But, this is little bit confusing when we need to send message, which is use HTTP GET to do that. But overall, it's a great API for me.

## How to use?

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

## ToDo
- [ ] Unit Testing
- [ ] Error Message & Response




