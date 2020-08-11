package main

import (
	"fmt"
	"reflect"
)

type TestStruc struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	s := TestStruc{
		Id:   12222,
		Name: "fffffff",
		Age:  1,
	}

	//v := reflect.ValueOf(s)
	//fmt.Println(v.FieldByName("Id").String())

	typ := reflect.TypeOf(s)

	field, err := typ.FieldByName("Id")
	fmt.Println(err)
	fmt.Println(field.Name)


	v := reflect.ValueOf(s)
	fmt.Println(v.FieldByName("Id").Int())

}
