package main

import (
        "log"
        "gopkg.in/ldap.v2"
        "crypto/tls"
        "crypto/x509"
        "io/ioutil" 
	"fmt"
)

var server_rootca string
var uri string
var server_name string
var bind_user string
var bind_pass string

func main() {

	var l  *ldap.Conn
	var conf *tls.Config

	// Read Input

	fmt.Print("Root CA Cert Path: ")
    	fmt.Scanln(&server_rootca)

	fmt.Print("URI (host:port): ")
	fmt.Scanln(&uri)

        fmt.Print("TLS Server Name: ")
        fmt.Scanln(&server_name)

        fmt.Print("Bind User: ")
        fmt.Scanln(&bind_user)

        fmt.Print("Bind Password: ")
        fmt.Scanln(&bind_pass)


	//Load CA Cert

	cert, err := ioutil.ReadFile(server_rootca)
	if err != nil {
		log.Fatalf("Couldn't load file", err)
	}
        
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(cert)

	// Connect

	conf = &tls.Config{
		RootCAs: certPool,
		InsecureSkipVerify: false,
		ServerName: server_name,
	}

	l, err = ldap.Dial("tcp", uri)

	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()

	err = l.StartTLS(conf)
	
	if err != nil {
		log.Fatal(err)
	}

	err = l.Bind(bind_user, bind_pass)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("OK: Connectivity verified\n")

}
