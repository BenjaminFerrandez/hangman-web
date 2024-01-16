package main

import "net/http"
import "fmt"

var port = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "contactez moi")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/contact", Contact)
	http.ListenAndServe(port, nil)
}
