package main

import (
	// "../smtpd"
	"bytes"
	"flag"
	"log"
	"net"
	"net/mail"
	"os"
	"strings"

	// load flag

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
func ListenAndServe(addr string, handler smtpd.Handler, authHandler smtpd.AuthHandler, appname string) error {
    srv := &smtpd.Server{
        Addr:        addr,
        Handler:     handler,
        Appname:     appname,
        Hostname:    "",
        AuthHandler: authHandler,
        AuthRequired: true,
    }
    return srv.ListenAndServe()
}

func main() {
	// read serverip, port, cert,key file, appname from args 
	var serverip string
	var port string
	var tls string
	var cert string
	var key string
	var appname string
	var enbaletls bool
	

	flag.StringVar(&serverip, "serverip", "10.223.40.21", "server ip")
	flag.StringVar(&port, "port", "2525", "server port")
	flag.StringVar(&tls, "tls", "6525", "server port")
	flag.StringVar(&cert, "cert", "server.pem", "cert file")
	flag.StringVar(&key, "key", "key.pem", "key file")
	flag.StringVar(&appname, "appname", "MyServerApp", "appname")
	flag.BoolVar(&enbaletls, "enbaletls", false, "enable tls")
	flag.Parse()



	log.Printf("serverip:%s, port:%s, cert:%s, key:%s, appname:%s", serverip, port, cert, key, appname)

// if cert and key file not exist, print error and exit

	if _, err := os.Stat(cert); os.IsNotExist(err) {
		log.Printf("cert file %s not exist", cert)
		return
	}
	if _, err := os.Stat(key); os.IsNotExist(err) {
		log.Printf("key file %s not exist", key)
		return
	}
	addr := serverip + ":" + port
	addrTls := serverip + ":" + tls

	smtpd.Debug = true
	if enbaletls {
    	go smtpd.ListenAndServeTLS(addrTls, cert, key, mailHandler, "tls " + appname, "", authHandler)
	} 
	
	go ListenAndServe(addr, mailHandler, authHandler, appname)
	

	select {}

}
