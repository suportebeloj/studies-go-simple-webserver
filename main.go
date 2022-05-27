package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("server runnig in http://127.0.0.1:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Invalid form %v", err)
		return
	}

	fmt.Fprintf(w, "POST request success")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprint(w, "hellow")
}
