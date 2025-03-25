# Passing a slice by reference

There are two ways to pass a slice to functions: by value and by pointer.

## 1. Pass by value

The slice is passed to the function by value, i.e. a copy of the original slice is created (length, capacity, and pointer to the underlying array). Changes inside the function do not affect the original slice outside, since the local variable exists only inside the function.

## 2. Pointer pass

The pointer to the slice is passed to the function, which allows the slice header itself (length, capacity, and pointer to the underlying array) to be modified. Changes inside the function affect the original slice outside, since the pointer refers to the same original slice.

## How to use?

If we are sure that reassigning a slice inside a function should not affect the original slice, then we should use value-based transfer. If it is important for the original slice to track all changes made to it, including slice replacement, then we should use pointer-based transfer.

It is also necessary to take into account that inside the function there may be functions such as `append`, which recreate the slice inside, and can lead to unexpected behavior if the slice is not passed by pointer.

## The `append` and `copy` method

The `append` function adds elements to the end of a slice. If the slice does not have enough capacity, a new underlying array is created. This can cause unexpected behavior. You may not see the changed values ​​outside the function.

```go
func append(slice []T, data ...T) []T {
    initialLength := len(slice)
    finalLength := m + len(data)
    if finalLength > cap(slice) { // if necessary, reallocate
        // allocate double what's needed, for future growth.
        newSlice := make([]byte, (n+1)*2)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0:finalLength]
    copy(slice[initialLength:finalLength], data)
    return slice
}
```

The `copy` function does not recreate the slice, but requires that we define a slice of sufficient capacity. No unexpected behavior can occur here.

```go
func copy(to []T, from []T) {
	for i := range from {
		to[i] = from[i]
	}
}
```

# How to use

## Run project

You can launch the project from the `cmd/basics/slices/reference-transfer` folder by the command:

```cmd
go run main.go
```

## Use entities

To use it in another project, you must do the following:

1. Import dependencies:

```go
import (
	"github.com/Mr-Filatik/go-knowledge-base/internal/basics/slices/modifier"
)
```

2. Use in code:

```go
modifier.ModifySliceDirectly(originalSlice, func(x int) int {
	return x * 2
})
modifier.ModifySliceReference(&originalSlice, func(x int) int {
	return x * 2
})
```

# Links

* [About Arrays and Slices in Go](https://habr.com/ru/articles/739754/ "Article on Habr")
* [Golang Tips: Why Slice Pointers Are Useful and How Ignoring Them Can Lead to Tricky Bugs](https://habr.com/ru/articles/525940/ "Article on Habr")
