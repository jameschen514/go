package main

import (
    "database/sql"
    "fmt"
    "html/template"
    "net/http"
    _ "github.com/go-sql-driver/mysql"
)

type Photo struct {
    Id      int
    Name    string
    Caption string
    FileString string
    Date    string
    Owner   string
}

type LoadPage struct {
    Pictures     []Photo
}

func dbConn() (db *sql.DB) {
    db, err := sql.Open("mysql", "root:root@tcp(192.168.1.195:8089/personalgo")
    if err != nil {
        fmt.Printf("Error: unable to connect to database")
    }
    return db
}

func process(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("templates/photos.html")
    if err != nil {
        fmt.Printf("Error: unable to read template.")
    }
    
    db := dbConn()
    selDB, err := db.Query("SELECT * FROM Photo ORDER BY id DESC")
    if err != nil {
        fmt.Printf("Error: unable to execute query")
    }

    var photos []Photo
    var photo Photo

    for selDB.Next() {
        /* fetch and define photo */
        var id int
        var name, caption, filestring, date, owner string

        err = selDB.Scan(&id, &name, &caption, &filestring, &date, &owner)
        if err != nil {
            fmt.Printf("Error: unable to load variables")
        }

        photo.Id = id
        photo.Name = name
        photo.Caption = caption
        photo.FileString = filestring
        photo.Date = date
        photo.Owner = owner

        photos = append(photos, photo)
    }

    t.Execute(w, photos)
    defer db.Close()
}

func main() {
    server := http.Server {
        Addr: "127.0.0.1:8000",
    }

    http.HandleFunc("/gallery", process)

    server.ListenAndServe()
}
