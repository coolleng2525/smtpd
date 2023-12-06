package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/mail"
	"strings"

	// "../smtpd"
	"github.com/mhale/smtpd"
)



func mailHandler(origin net.Addr, from string, to []string, data []byte) error {
    msg, _ := mail.ReadMessage(bytes.NewReader(data))
    subject := msg.Header.Get("Subject")
    log.Printf("Received mail from %s for %s with subject %s", from, to[0], subject)
    return nil
}

func authHandler(remoteAddr net.Addr, mechanism string, username []byte, password []byte, shared []byte) (bool, error) {
	log.Printf("Authenticating %s %s %s %s", remoteAddr, mechanism, username, password)
	// remoteadd start with 10.223.40 return true
	a := strings.Split(remoteAddr.String(), ".")
	if a[0] == "10" && a[1] == "223" && a[2] == "40" {
		log.Printf("It's from 10.223.40.x, return true")
		return true, nil
	}
	if string(username) == "aa" && string(password) == "b729b68ff568e2e7bd9e831ba7982c77" {
		log.Printf("It's default username:%s , return true", username)
		return true, nil
	}
	
	return false, nil
}

func main() {
	smtpd.Debug = true
	fmt.Println("Starting server on port 2525")
    smtpd.ListenAndServe("10.223.40.21:2525", mailHandler, "MyServerApp", "", authHandler)
}