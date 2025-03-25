package modifier

import "fmt"

func ModifySliceDirectly(slice []int, modFunc func(int) int) {
	fmt.Println("Inside modifyDirectly before re-creation:", slice)

	// Somewhere in the code the original slice is replaced with a new one
	newSlice := make([]int, len(slice), cap(slice))
	copy(newSlice, slice)
	for ind := range newSlice {
		newSlice[ind] = modFunc(newSlice[ind])
	}
	slice = newSlice // Replacement

	fmt.Println("Inside modifyDirectly after re-creation:", slice)
}

func ModifySliceReference(slicePtr *[]int, modFunc func(int) int) {
	fmt.Println("Inside modifySliceReference before re-creation:", *slicePtr)

	// Somewhere in the code, the original slice is replaced with a new one via a pointer
	newSlice := make([]int, len(*slicePtr), cap(*slicePtr))
	copy(newSlice, (*slicePtr))
	for ind := range newSlice {
		newSlice[ind] = modFunc(newSlice[ind])
	}
	*slicePtr = newSlice // Replacement by pointer

	fmt.Println("Inside modifySliceReference after re-creation:", *slicePtr)
}
