package main

import (
    "net/http"
    "html/template"
    "os"
    "log"
)

func view(w http.ResponseWriter, r *http.Request) {
    htmlFile, err := os.Open("data.html")
    if err != nil {
        log.Fatal("Error opening file")
    }

    t, _ := template.ParseFiles("loop.html")
    content := htmlFile
    t.Execute(w, content)
}

func main() {
    server := http.Server {
        Addr: "localhost:8080",
    }
    http.HandleFunc("/view", view)
    server.ListenAndServe()
}
