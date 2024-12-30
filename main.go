package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not Found" , http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method Is Not Supported" , http.StatusNotFound)
		return
	}

	fmt.Fprintf(w,"Hello Hii");
}

func formHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        return
    }

    if err := r.ParseForm(); err != nil {
        http.Error(w, fmt.Sprintf("ParseForm() error: %v", err), http.StatusBadRequest)
        return
    }

    name := r.FormValue("name")
    address := r.FormValue("address")

    if name == "" || address == "" {
        http.Error(w, "Name and Address are required", http.StatusBadRequest)
        return
    }

    fmt.Fprintf(w, "POST Request Successful\n")
    fmt.Fprintf(w, "Name = %s\n", name)
    fmt.Fprintf(w, "Address = %s\n", address)
}


func main() {
	// fmt.Println("Hello Go............")
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/" , fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello" , helloHandler)

	fmt.Printf("Starting Server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
 
}