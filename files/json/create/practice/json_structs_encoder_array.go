package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type Post struct {
    Id          int         `json:"id"`
    Content     string      `json:"content"`
    Author      Author      `json:"author"`
    Comments    []Comment   `json:"comments"`
}

type Author struct {
    Id      int     `json:"id"`
    Name    string  `json:"name"`
}

type Comment struct {
    Id      int     `json:"id"`
    Content string  `json:"content"`
    Author  string  `json:"author"`
}

func main() {
    post := []Post { Post {
        Id: 1,
        Content: "Hello World!",
        Author: Author {
            Id: 2,
            Name: "Sau Sheong",
        },
        Comments: []Comment {
            Comment {
                Id: 3,
                Content: "Have a great day!",
                Author: "Adam",
            },
            Comment {
                Id: 4,
                Content: "How are you today?",
                Author: "Betty",
            },
        },
    }, Post {
        Id: 2,
        Content: "Hello There World!",
        Author: Author {
            Id: 3,
            Name: "Baboom",
        },
        Comments: []Comment {
            Comment {
                Id: 5,
                Content: "What's for dinner?",
                Author: "James",
            },
            Comment {
                Id: 6,
                Content: "What's for lunch?",
                Author: "Erika",
            },
        },
}}
   
    var newPost = Post { 
                    Id: 3, 
                    Content: "How do you do world!", 
                    Author: Author {
                        Id: 4, Name: "Ahwooga",
                    },
                    Comments: []Comment {
                        Comment {
                            Id: 7,
                            Content: "What's for breakfast",
                            Author: "James",
                        },
                    } }


    post = append(post, newPost)

    jsonFile, err := os.Create("post_encoder_append.json")
    if err != nil {
        fmt.Println("Error creating JSON file: ", err)
        return
    }

    encoder := json.NewEncoder(jsonFile)
    err = encoder.Encode(&post)
    if err != nil {
        fmt.Println("Error encoding JSON to file: ", err)
        return
    }
}
