package main

import (
	"fmt"	
	"encoding/json"	// Encoding and Decoding Package
    "reflect"
)

type Employee struct {
    Id          int
    Name        string
    Jobtitle    string
    Phone       Phone
    Email       string
}

type Phone  struct {
    Home    string
    Office  string
}

var JSON = `[
    {
        "Id": 1,
	    "Name":"Mark Taylor",
	    "Jobtitle":"Software Developer",
	    "Phone":{
		    "Home":"123-466-799",
		    "Office":"564-987-654"
	    },
	    "Email":"markt@gmail.com"
    },
    {

        "Id": 2,
	    "Name":"Mark Taylor",
	    "Jobtitle":"Software Developer",
	    "Phone":{
		    "Home":"123-466-799",
		    "Office":"564-987-654"
	    },
        "Email":"asdf@asdf.com"
    }
]`

func main() {
	// Unmarshal the JSON string into info map variable.
	//var info map[string]interface{}
    fmt.Printf("JSONType: %s", reflect.TypeOf(JSON))
    var employees []Employee
    err := json.Unmarshal([]byte(JSON),&employees)
    if err != nil {
        fmt.Println("Error: ", err)
    }

    fmt.Printf("%+v", employees)
    fmt.Printf("%s", employees)
    //fmt.Printf("id: %d", employees[0].Id)
    //fmt.Printf("id: %d", employees[0].Id)
    // Print the output from info map.
	//fmt.Println(info)
	//fmt.Println(info["jobtitle"])
	//fmt.Println(info["email"])
	//fmt.Println(info["phone"].(map[string]interface{})["home"])	
	//fmt.Println(info["phone"].(map[string]interface{})["office"])
}


