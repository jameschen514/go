/*********************
*
* By: James Chen
* Contact: james.chen@seethisthing.com
*
*********************/

package main

import (
    "fmt"
    //"reflect"
    "strconv"
)


func bool2str(value bool) string {
    
    result := fmt.Sprintf("%s", value)

    switch(value) {
        case true:
            result = "1"
        case false:
            result = "0"
        default: 
            result = "-1"
    }
    
    return result
}

func str2bool(value string) bool {

    result, err := strconv.ParseBool(value)

    if err != nil {
        fmt.Println("Error: unable to parse input")
    }

    return result
}

func main() {
    var value bool = true
    fmt.Println("Value: ", bool2str(value))
   
    var word string = "0"
    fmt.Println("Value2: ", str2bool(word))
}
