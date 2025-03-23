package main

import (
	"fmt"
)

func main() {
	// 1. Declaring a fixed length array
	fmt.Println("----- Declaring a fixed length array -----")
	var arr1 [5]int // An array of 5 integers, all elements initialized to zeros
	fmt.Println("arr1:", arr1)

	// 2. Initializing an array when declaring it
	fmt.Println("\n----- Initializing an array when declaring it -----")
	arr2 := [5]int{1, 2, 3, 4, 5} // Explicit initialization of all elements
	fmt.Println("arr2:", arr2)

	// 3. Initializing an array with a length skip
	fmt.Println("\n----- Initializing an array with a length skip -----")
	arr3 := [...]int{10, 20, 30, 40, 50} // The compiler determines the length automatically
	fmt.Println("arr3:", arr3)

	// 4. Accessing array elements
	fmt.Println("\n----- Accessing array elements -----")
	fmt.Println("The first element of arr3:", arr3[0]) // Indexing starts from 0
	fmt.Println("The last element of arr3:", arr3[len(arr3)-1])

	// 5. Changing array elements
	fmt.Println("\n----- Changing array elements -----")
	arr3[0] = 100
	fmt.Println("arr3 after changing the first element:", arr3)

	// 6. Looping through an array using a for loop
	fmt.Println("\n----- Looping through an array using a for loop -----")
	fmt.Print("Looping through the elements of arr3 using a for loop: ")
	for i := 0; i < len(arr3); i++ {
		fmt.Print(arr3[i], " ")
	}
	fmt.Println()

	// 7. Looping through array elements using for-range loop
	fmt.Println("\n----- Looping through array elements using for-range loop -----")
	fmt.Print("Looping through the elements of arr3 using a for-range loop: ")
	for index, value := range arr3 {
		fmt.Printf("arr3[%d]=%d ", index, value)
	}
	fmt.Println()

	// 8. Multidimensional arrays (matrices)
	fmt.Println("\n----- Multidimensional arrays (matrices) -----")
	var matrix [3][3]int
	matrix[0] = [3]int{1, 2, 3}
	matrix[1] = [3]int{4, 5, 6}
	matrix[2] = [3]int{7, 8, 9}
	fmt.Println("Matrix:")
	for _, row := range matrix {
		fmt.Println(row)
	}

	// 9. Copying arrays
	fmt.Println("\n----- Copying arrays -----")
	arr4 := arr3 // Array copying occurs by value, not by reference
	arr4[0] = 999
	fmt.Println("arr3 after changing arr4:", arr3) // arr3 has not changed
	fmt.Println("arr4:", arr4)

	// 10. Comparing arrays
	fmt.Println("\n----- Comparing arrays -----")
	arr5 := [5]int{100, 20, 30, 40, 50}
	if arr3 == arr4 {
		fmt.Println("arr3 and arr4 are equal")
	} else {
		fmt.Println("arr3 and arr4 are not equal")
	}
	if arr3 == arr5 {
		fmt.Println("arr3 and arr5 are equal")
	} else {
		fmt.Println("arr3 and arr5 are not equal")
	}
}
