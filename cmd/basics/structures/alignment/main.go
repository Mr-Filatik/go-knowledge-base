package main

import (
	"github.com/Mr-Filatik/go-knowledge-base/internal/basics/structures/alignment/printer"
	"github.com/Mr-Filatik/go-knowledge-base/internal/basics/structures/alignment/scanner"
)

type MyStruct struct {
	Field1 int8
	Field2 int8
	Field3 int16
	Field4 int32
	Field5 int64
	Field6 int16
}

func main() {
	ms := MyStruct{}

	scan := scanner.ScanStructInfo(ms, "Scan My Struct")
	printer.PrintStructInfo(scan)
}
