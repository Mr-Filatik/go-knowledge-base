# Increasing slice capacity

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

You can launch the project from the `cmd/basics/slices/capacity-increase` folder by the command:

```cmd
go run main.go
```

## Use entities

To use it in another project, you must do the following:

1. Import dependencies:

```go
import (
	"github.com/Mr-Filatik/go-knowledge-base/internal/basics/slices/capacity-increase/printer"
	"github.com/Mr-Filatik/go-knowledge-base/internal/basics/slices/capacity-increase/scanner"
)
```

2. Use in code:

```go
capacities := scanner.ScanSliceCapacityIncrease(4, 20000)
printer.PrintSliceCapacityIncrease(capacities)
```

# Links

* [About Arrays and Slices in Go](https://habr.com/ru/articles/739754/ "Article on Habr")
* [Golang Tips: Why Slice Pointers Are Useful and How Ignoring Them Can Lead to Tricky Bugs](https://habr.com/ru/articles/525940/ "Article on Habr")
