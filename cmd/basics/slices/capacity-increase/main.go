package main

import (
	"github.com/Mr-Filatik/go-knowledge-base/internal/basics/slices/capacity-increase/printer"
	"github.com/Mr-Filatik/go-knowledge-base/internal/basics/slices/capacity-increase/scanner"
)

func main() {
	capacities := scanner.ScanSliceCapacityIncrease(4, 10000)
	printer.PrintSliceCapacityIncrease(capacities)
}
