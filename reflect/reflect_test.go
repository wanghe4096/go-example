package main

import (
	"reflect"
	"testing"
)

func TestCase(t *testing.T) {
	s := TestStruc{
		Id:   1,
		Name: "fffffff",
		Age:  1,
	}

	typ := reflect.TypeOf(s)
	t.Log(typ.Name())

}
