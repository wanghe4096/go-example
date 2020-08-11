package main

import (
	"container/list"
	"reflect"
)

func main() {
	l := list.New()
	for i := 0; i < 10; i++ {
		l.
	}

	for i := 0; i < 10; i++ {
		println(l.Front().Value.(int))
		println(reflect.TypeOf(l.Front().Value).Name())

	}
}
