package main

import (
    "fmt"
    "encoding/json"
)

type Person struct {
	Name string
	Adress string
}

func main(){
	var name string
	var adress string
	fmt.Println("Please enter a name:")
	fmt.Scanln(&name)
	fmt.Println("Please enter a adress:")
	fmt.Scanln(&adress)

    	in := Person{Name: name, Adress: adress}

    	var inInterface map[string]interface{}
    	inrec, err := json.Marshal(in)
    	json.Unmarshal(inrec, &inInterface)
		
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(inrec))
	}
}


