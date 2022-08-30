package main

import (
    "encoding/xml"
    "fmt"
    "io"
    "os"
)

type Post struct {
    XMLName     xml.Name    `xml:"post"`
    Id          string      `xml:"id, attr"`
    Content     string      `xml:"content"`
    Author      Author      `xml:"author"`
    Comments    []Comment   `xml:"comments>comment"`
}

type Author struct {
    Id      string      `xml:"id, attr"`
    Name    string      `xml:", chardata"`
}

type Comment struct {
    Id      string      `xml:"id, attr"`
    Content string      `xml:"content"`
    Author  Author      `xml:"author"`
}

func main() {

    xmlFile, err := os.Open("nestedelement.xml")
    if err != nil {
        fmt.Println("Error opening XML file: ", err)
        return
    }
    defer xmlFile.Close()

    decoder := xml.NewDecoder(xmlFile)
    for {
        t, err := decoder.Token()
        if err == io.EOF {
            break
        }

        if err != nil {
            fmt.Println("Error decoding XML into tokens: ", err)
            return
        }

        switch se := t.(type) {
            case xml.StartElement:
                if se.Name.Local == "post" {
                    fmt.Printf("Element: %s \n", se.Attr[0].Value)
                }
                fmt.Printf("Element: %s \n", se.Name)
                if se.Name.Local == "comment" {
                    var comment Comment
                    decoder.DecodeElement(&comment, &se)
                }
            case xml.CharData:
                fmt.Printf("Chardata: %s \n", string([]byte(xml.CharData(se))))
        }

    }

}
