package main

import (
    "net/http"
    "html/template"
    "time"
)

func process(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("main.html")
    date := time.Now()

    t.Execute(w, date)
}

func main() {
    server := http.Server {
        Addr: "127.0.0.1:8080",
    }

    http.HandleFunc("/process", process)

    server.ListenAndServe()
}
