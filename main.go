package main

import (
	"fmt"
	"net/http"
)

func main() {
	
	http.HandleFunc("/", handler)

	http.ListenAndServeTLS(":8080", "./sshKey/cert.pem", "./sshKey/key.pem", nil)

}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World")
}
