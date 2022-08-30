package main

import (
    "encoding/base64"
    "fmt"
    "io/ioutil"
    //"log"
    "net/http"
    "html/template"
)

type Image struct {
    ImageString     string
}

func process(w http.ResponseWriter, r *http.Request) {

    t, _ := template.ParseFiles("index.html")


    img, err := ioutil.ReadFile("toyota-supra1.jpg")
    if err != nil {
        fmt.Printf("Error: unable to read file")
    }

    encoded := base64.StdEncoding.EncodeToString([]byte(img))

    fmt.Printf("BASE64: %s", encoded)

    image := Image {}
    image.ImageString = encoded
    
    fmt.Printf("%s", encoded)

    t.ExecuteTemplate(w, "layout", image)    
}

func main() {

    server := http.Server {
        Addr: "127.0.0.1:8080",
    }

    http.HandleFunc("/process", process)

    

    server.ListenAndServe()
}
