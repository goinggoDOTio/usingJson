package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	//using json marshall
	fmt.Println("---JSON MARSHALL---")
	p1 := Person{"James", "Bond", 20, 007}
	bs, err := json.Marshal(p1)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))

	// using json unmarshalfmt.Println("---USING JSON UNMARSHALL---")
	fmt.Println("---JSON UNMARSHALL---")
	var p2 Person
	fmt.Println(p2.First)
	fmt.Println(p2.Last)
	fmt.Println(p2.Age)

	bs2 := []byte(`{"First":"James", "Last":"Bond","Age":20}`)
	json.Unmarshal(bs2, &p2)
	fmt.Println("-----------------------")
	fmt.Println(p2.First)
	fmt.Println(p2.Last)
	fmt.Println(p2.Age)

	// using json en$coding
	fmt.Println("---JSON ENCODING---")
	p3 := Person{"James", "Bond", 20, 007}
	json.NewEncoder(os.Stdout).Encode(p3)

	// using json decoding by simulating a json data comming in
	fmt.Println("---JSON DECODING---")
	var p4 Person

	fakeJsonStream := strings.NewReader(`{"First":"James", "Last":"Bond","Age":20}`)
	json.NewDecoder(fakeJsonStream).Decode(&p4)

	fmt.Println(p4.First)
	fmt.Println(p4.Last)
	fmt.Println(p4.Age)
}
