package main

import (
    "log"
    "net/http"
)

func main() {
    fs := http.FileServer(http.Dir("./asset/images"))
    http.Handle("/asset/images/", http.StripPrefix("/asset/images", fs))
   
    js := http.FileServer(http.Dir("./asset/scripts/"))
    http.Handle("/asset/scripts/", http.StripPrefix("/asset/scripts", js))

    css := http.FileServer(http.Dir("./asset/styles/"))
    http.Handle("/asset/styles/", http.StripPrefix("/asset/styles", css))
    

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
