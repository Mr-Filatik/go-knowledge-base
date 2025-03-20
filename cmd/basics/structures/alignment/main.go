package main

import (
	"github.com/Mr-Filatik/go-knowledge-base/internal/basics/structures/alignment/printer"
	"github.com/Mr-Filatik/go-knowledge-base/internal/basics/structures/alignment/scanner"
)

type MyStruct struct {
	aaa  int8
	aaa2 int8
	ccc  int16
	ddd  int32
	bbb  int64
	ggg  int16
}

func main() {
	ms := MyStruct{
		aaa:  1,
		bbb:  2,
		ccc:  3,
		ddd:  4,
		aaa2: 5,
		ggg:  6,
	}

	scan := scanner.ScanStructInfo(ms, "example one")
	printer.PrintStructInfo(scan)
}
