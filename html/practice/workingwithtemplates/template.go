package main

import (
    "fmt"
    "html/template"
    "net/http"
)

type Name struct {
    FirstName       string
    LastName        string
}

type Feed struct {
    Title   string
    Text    string
}

type LoadPage struct {
    Heading     Name
    NewsFeed    []Feed
}

func process(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("index.html")
    if err != nil {
        fmt.Printf("Error: unable to open file")
    }

    //n := &Name { FirstName: "James", LastName: "Chen" }

    load := &LoadPage {
            Heading: Name { FirstName: "James", LastName: "Chen"}, 
            NewsFeed : []Feed { Feed { Title : "My first load page", Text: "^OMG it's so easy"}, },
    }
    
    t.Execute(w, load)
}

func main() {
    server := http.Server {
        Addr: "127.0.0.1:8080",
    }

    http.HandleFunc("/", process)

    server.ListenAndServe()
}
