package main

import (
	"fmt"
)

func main() {
	// 1. Creating a slice using make
	fmt.Println("----- Creating a slice using make -----")
	slice1 := make([]int, 5) // Slice of 5 elements, all initialized to zeros
	fmt.Println("slice1:", slice1)

	// 2. Creating a slice with explicit initialization
	fmt.Println("\n----- Creating a slice with explicit initialization -----")
	slice2 := []int{1, 2, 3, 4, 5} // Explicitly specifying values
	fmt.Println("slice2:", slice2)

	// 3. Create a slice with length and capacity specified
	fmt.Println("\n----- Create a slice with length and capacity specified -----")
	slice3 := make([]int, 3, 5) // Length = 3, capacity = 5
	fmt.Println("slice3:", slice3)
	fmt.Println("Length slice3:", len(slice3), "Capacity slice3:", cap(slice3))

	// 4. Accessing slice elements
	fmt.Println("\n----- Accessing slice elements -----")
	fmt.Println("First element of slice2:", slice2[0])
	fmt.Println("Last element of slice2:", slice2[len(slice2)-1])

	// 5. Changing slice elements
	fmt.Println("\n----- Changing slice elements -----")
	slice2[0] = 100
	fmt.Println("slice2 after changing the first element:", slice2)

	// 6. Looping through a slice using a for loop
	fmt.Println("\n----- Looping through a slice using a for loop -----")
	fmt.Print("Looping through the elements of slice2 using a for loop: ")
	for i := 0; i < len(slice2); i++ {
		fmt.Print(slice2[i], " ")
	}
	fmt.Println()

	// 7. Looping through a slice using the for-range loop
	fmt.Println("\n----- Looping through a slice using the for-range loop -----")
	fmt.Print("Slice2 elements via for-range:")
	for index, value := range slice2 {
		fmt.Printf("slice2[%d]=%d ", index, value)
	}
	fmt.Println()

	// 8. Adding Elements to a slice with append
	fmt.Println("\n----- Adding Elements to a slice with append -----")
	slice2 = append(slice2, 6, 7, 8)
	fmt.Println("slice2 after adding elements:", slice2)

	// 9. Removing an element from a slice
	fmt.Println("\n----- Removing an element from a slice -----")
	indexToRemove := 2
	slice2 = append(slice2[:indexToRemove], slice2[indexToRemove+1:]...)
	fmt.Println("slice2 after removing the element with index 2:", slice2)

	// 10. Creating a subslice (slice slicing)
	fmt.Println("\n----- Creating a subslice (slice slicing) -----")
	subSlice := slice2[1:4] // Elements with indices 1, 2, 3
	fmt.Println("subSlice:", subSlice)
	slice2[2] = -1
	fmt.Println("subSlice after slice2 change:", subSlice)

	// 11. Copying slices
	fmt.Println("\n----- Copying slices -----")
	slice4 := make([]int, len(slice2)) // It is necessary to create a cut of sufficient length
	copy(slice4, slice2)               // Copying elements of slice2 to slice4
	fmt.Println("slice4 (copy of slice2):", slice4)
	slice2[2] = -2 // Copies of all elements are created, the slices after copying are not connected
	fmt.Println("slice4 (copy of slice2) after changing slice2:", slice4)

	// 12. Multidimensional slices
	fmt.Println("\n----- Multidimensional slices -----")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Matrix (multidimensional slice):")
	for _, row := range matrix {
		fmt.Println(row)
	}
}
