package main

import (
    "fmt"
    "net/http"
    "os"
    "log"
)

func process(w http.ResponseWriter, r *http.Request) {
    file, err := os.Open("file.go")
    if err != nil {
        log.Fatal(err)
    }

    data := make([]byte, 100)
    count, err := file.Read(data)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("read %d bytes: %q\n", count, data[:count])
}

func main() {
    server := http.Server {
        Addr: "127.0.0.1:8080",
    }

    http.HandleFunc("/process", process)

    server.ListenAndServe()
}
