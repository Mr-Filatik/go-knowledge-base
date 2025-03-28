package main

import (
	"fmt"
)

func main() {
	var interfaceVar MyInterface
	Compare("MyInterface (nil)", interfaceVar)

	var structNil *MyStruct
	interfaceVar = structNil
	Compare("MyInterface (*MyStruct = nil)", interfaceVar)

	var structNotNil *MyStruct = &MyStruct{}
	interfaceVar = structNotNil
	Compare("MyInterface (*MyStruct != nil)", interfaceVar)
}

type MyInterface interface {
	MyFunc()
}

type MyStruct struct {
}

func (e *MyStruct) MyFunc() {
	fmt.Println("MyStruct.IntfFunc from MyInterface")
}

func Compare(name string, v any) {
	if v == nil {
		fmt.Println(name, "is nil")
	} else {
		fmt.Println(name, "is not nil")
	}
}
