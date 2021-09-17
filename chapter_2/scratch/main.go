package main

import (
	"encoding/json"
	"fmt"
)

type thing struct {
	FieldOne string `json:"field1"`
}

type thing2 struct {
	Fieldn string
	Field1 string
	Field2 string
}

func main() {

	t2 := thing2 {
		Fieldn: "var",
		Field1: "one",
		Field2: "two",
	}

	fmt.Println(t2)

	t1 := thing{}

	b, err := json.Marshal(t2)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))

	json.Unmarshal(b, &t1)

	fmt.Println(t1)
}
