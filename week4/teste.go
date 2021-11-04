package main

import (
    "fmt"
    "encoding/json"
)

type MyData struct {
    One   int
    Two   string
    Three int
}

func main() {   
    in := &MyData{One: 1, Two: "second"}

    var inInterface map[string]interface{}
    inrec, _ := json.Marshal(in)
    json.Unmarshal(inrec, &inInterface)

    // iterate through inrecs
    for field, val := range inInterface {
            fmt.Println("KV Pair: ", field, val)
    }
}

