package study

import (
	"testing"
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string	`json:"name"`
	Age int32	`json:"age"`
	Height int32
}

func TestJson(t *testing.T){
	person := &Person{"taiziwan" , 18 , 178}

	jsonStr,_ := json.Marshal(person)
	fmt.Println(string(jsonStr))

	var person_decode  Person

	_ = json.Unmarshal(jsonStr,&person_decode)

	fmt.Println(person_decode.Age)
}
