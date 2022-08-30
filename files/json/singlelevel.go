package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    //"reflect"
)

type Post struct {
    Id      int         `json:"id"`
    Content string      `json:"content"`
    Author  string      `json:"author"`
    Comments string     `json:"comments"`
}

func main() {
    jsonFile, err := os.Open("data.json")
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

    //fmt.Println(string(jsonData))

    var post Post
    json.Unmarshal(jsonData, &post)
    fmt.Println(post)
}
