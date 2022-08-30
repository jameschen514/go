/********************
*
* Written by: James Chen
* Contact: jamescychen@gmail.com
*
********************/


package main

import (
    "fmt"
    "flag"
    "os"
    "io/ioutil"
)

func main() {
    fmt.Printf("Usage: mvformove -o [origin path] -d [destination path]\n")
    var origin = flag.String("o", "",  "ex. mvformove -o '/home/username/desktop'")
    var destination = flag.String("d", "", "ex. mvformove -d '/home/username/wwwroot'")
    flag.Parse()
    
    if *origin == "" || *destination == "" {
        fmt.Printf("Error: Origin and/or destination is not defined\n")
    } else {
        files, err := ioutil.ReadDir(*origin)
        if err != nil {
            fmt.Printf("Error: unable to read files defined in origin path")
        }

        for i := range files {
            fmt.Printf("a: %s b: %s", files[i].Name(), *destination)
            var newpath string = *destination + "/" + files[i].Name()
            fmt.Printf("new path: %s\n", newpath)
            os.Rename(files[i].Name(), newpath)
        }

    } 
     
    //fmt.Printf("source: %s\ndestionation: %s\n", *origin, *destination)
}
