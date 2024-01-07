package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ResponseRequest struct{
    Authorization string `json:"authorization"`
    ContentType string `json:"contenttype"`
    ContentLength int `json:"contentlength"`
    Method string `json:"method"`
    Body string `json:"body"`
    URL string `json:"url"`
}

func echoRequest(w http.ResponseWriter, r *http.Request){
    rr := ResponseRequest{
        Authorization: r.Header.Get("Authorization"),
        ContentType: r.Header.Get("Content-Type"),
        Method: r.Method,
        URL: r.URL.String(),
    }
    // get content-length
    contentLengthRaw := r.Header.Get("Content-Length")
    if contentLengthRaw == ""{
        rr.ContentLength = -1
    }else{
        contentLength, err := strconv.Atoi(contentLengthRaw)
        if err != nil{
            fmt.Fprintf(w, "error: convert content-length.")
            fmt.Printf("%v", err)
            return
        }
        rr.ContentLength = contentLength
    }

    // get request body
    bodyBytes := make([]byte, rr.ContentLength)
    r.Body.Read(bodyBytes)
    rr.Body = string(bodyBytes)

    resp, err := json.Marshal(rr)
    if err != nil{
        fmt.Fprintf(w, "error: convert json error.")
        fmt.Printf("%v", err)
        return
    }

    outputResponseRequest(&rr)
    fmt.Fprintf(w, string(resp))
}

func main(){
    addr := fmt.Sprintf("%s:%d", "0.0.0.0", 8983)
    http.HandleFunc("/", echoRequest)

    fmt.Printf("echo-server listen: '%s'\n", addr)
    if err := http.ListenAndServe(addr, nil);err != nil{
        panic(err)
    }
}
