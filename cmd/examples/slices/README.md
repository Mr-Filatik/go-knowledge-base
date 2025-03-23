# Slices

Basic rules for working with slices:

- If you do not specify explicit values ​​for slice elements, all elements are initialized with the default value for the type (for example: 0 for int, etc.);
- Slice elements are indexed starting from the 0th element;
- A slice has a length (number of elements) and a capacity (maximum size of the base array). If the capacity is exceeded, a new base array of a larger size is created and the current elements are copied to it;
- The values ​​of slice elements can be changed by accessing them by index;
- To create a subslice, the initial index is inclusive, the final index is exclusive;
- To copy slices, the `copy` function is used, which creates a full copy of the slice (copies of all elements of the base array);
- To add elements to a slice, the `append` method is used (it can re-create the slice);
- To compare slices, all their elements must be compared.

A slice under the hood is a structure that contains a reference to the original array, its length, and its capacity:

```go
struct {
	array *[]T
	length int
	capacity int
}
```

When the capacity is exceeded, the base array is recreated with a larger array. This value is not fixed and depends on the size of the current capacity. The values ​​are given in the table (the values ​​may change in new versions of the language):

| Starting cap  | Growth factor |
| ------------- | ------------- |
| 256 | 2.0 |
| 512 | 1.63 |
| 1024 | 1.44 |
| 2048 | 1.35 |
| 4096 | 1.30 |

Doubling the capacity for small slices minimizes the number of memory allocations, which improves performance. However, for large slices, doubling the capacity would result in significant memory overhead, so a slower increase is used.

# How to use

## Run project

You can launch the project from the `cmd/examples/slices` folder by the command:

```cmd
go run main.go
```

# Links

* [About Arrays and Slices in Go](https://habr.com/ru/articles/739754/ "Article on Habr")
* [Golang Tips: Why Slice Pointers Are Useful and How Ignoring Them Can Lead to Tricky Bugs](https://habr.com/ru/articles/525940/ "Article on Habr")

Press **[Back](../README.md "Return to previous README.md")** to go to the previous `README.md` file.