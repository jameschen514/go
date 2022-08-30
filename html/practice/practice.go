package main

import (
    "fmt"
    "net/http"
    "html/template"
)

func process(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("practice.html")
    var text string 
    text = "hello there"
    fmt.Printf("____")
    t.Execute(w, text)
}

func main() {
    server := http.Server {
        Addr: "127.0.0.1:8080",
    }

    http.HandleFunc("/process", process)
    server.ListenAndServe()
}
