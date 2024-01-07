package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func outputResponseRequest(rr *ResponseRequest){
    fmt.Println("Request Recieved")
    fmt.Println("===================================")
    fmt.Printf("Authorization: %s\n", rr.Authorization)
    fmt.Printf("Content-Type: %s\n", rr.ContentType)
    fmt.Printf("Method: %s\n", rr.Method)
    fmt.Printf("URL: %s\n", rr.URL)

    fmt.Println("Body")
    fmt.Println("-----------------------------------")
    if rr.ContentType != "application/json"{
        fmt.Println(rr.Body)
    }else{
        var buf bytes.Buffer
        err := json.Indent(&buf, []byte(rr.Body), "", " ")
        if err != nil{
            panic(err)
        }
        fmt.Println(buf.String())
    }
    fmt.Println("-----------------------------------")

    fmt.Println("===================================")
    fmt.Println()
    fmt.Println()
}
