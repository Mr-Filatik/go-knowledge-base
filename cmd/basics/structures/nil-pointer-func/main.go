package main

import "fmt"

func main() {
	var structVar *MyStruct
	fmt.Println("nil.MyFuncSave()")
	structVar.MyFuncSave()
	fmt.Println("nil.MyFuncUnsave()")
	// structVar.MyFuncUnsave() // panic: runtime error: invalid memory address or nil pointer dereference

	structVar = &MyStruct{
		Value: "value",
	}
	fmt.Println("*MyStruct.MyFuncSave()")
	structVar.MyFuncSave()
	fmt.Println("*MyStruct.MyFuncUnsave()")
	structVar.MyFuncUnsave()
}

type MyStruct struct {
	Value string
}

func (e *MyStruct) MyFuncSave() {
	if e == nil {
		fmt.Println("I'm nil")
		return
	}
	fmt.Printf("I'm not nil, %s\n", e.Value)
}

func (e *MyStruct) MyFuncUnsave() {
	fmt.Printf("I'm not nil, %s\n", e.Value)
}
