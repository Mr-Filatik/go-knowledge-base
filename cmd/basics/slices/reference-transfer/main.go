package main

import (
	"fmt"

	"github.com/Mr-Filatik/go-knowledge-base/internal/basics/slices/modifier"
)

func main() {
	originalSlice := []int{1, 2, 3} // Original slice
	fmt.Println("Original slice:", originalSlice)

	fmt.Println()
	modifier.ModifySliceDirectly(originalSlice, func(x int) int {
		return x * 2
	})
	fmt.Println("Original slice after modifySliceDirectly:", originalSlice)

	fmt.Println()
	modifier.ModifySliceReference(&originalSlice, func(x int) int {
		return x * 2
	})
	fmt.Println("Original slice after modifySliceReference:", originalSlice)
}
