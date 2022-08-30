package main

import (
    "encoding/json"
    "os"
    "fmt"
    "io/ioutil"
)

type Story struct {
    Id          int
    Content     string
    Author      Author
    Comments    []Comments
}

type Author struct {
    Id          int
    Name        string
}

type Comments struct {
    Id      int
    Content string
    Author  string
}

func main() {
    // open file
    fh, err := os.Open("multiple.json")
    if err != nil {
        fmt.Printf("Error opening file.")
    }
    defer fh.Close()

    // Read file
    fhData, err := ioutil.ReadAll(fh)
    if err != nil {
        fmt.Println("Error reading json data: ", err)
        return
    }

    var stories []Story
    err = json.Unmarshal(fhData, &stories)
    if err != nil {
        fmt.Println("Error: ", err)
    }
    // print file data
    fmt.Printf("%+v", stories)
}
