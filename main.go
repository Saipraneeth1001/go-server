package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello Rudra Deva")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm();
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Fprintf(w, "post request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "name = %s\n", name)
	fmt.Fprintf(w, "address = %s\n", address)
}

func main() {

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting server at port 8080\n");
	err := http.ListenAndServe(":8080", nil); 
	if err != nil {
        log.Fatal(err)
    }



}