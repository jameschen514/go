package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

type Post struct {
    Id      int         `json:"id"`
    Content string      `json:"content"`
    Author  string      `json:"author"`
    Comments string     `json:"comments"`
}

func main() {
    jsonFile, err := os.Open("info.json")
    if err != nil {
        fmt.Println("Error opening JSON file: ", err)
        return
    }

    defer jsonFile.Close()

    jsonData, err := ioutil.ReadAll(jsonFile)
    if err != nil {
        fmt.Println("Error reading JSON data: ", err)
        return
    }

    var post Post
    json.Unmarshal(jsonData, &post)
    fmt.Println(post.Id)
}
