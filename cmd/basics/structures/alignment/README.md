# Structure alignment

Structure alignment in Go is done to optimize memory access. Here are the basic principles:

## 1. Field alignment

Each data type has a "natural" alignment equal to its size:
- Types of size 1 byte (e.g., `int8`, `uint8`) are aligned on a 1-byte boundary;
- Types of size 2 bytes (e.g., `int16`, `uint16`) are aligned on a 2-byte boundary;
- Types of size 4 bytes (e.g. `int32`, `float32`) are aligned on a 4 byte boundary;
- Types of size 8 bytes (e.g., `int64`, `float64`, pointers, slice, map, strings) are aligned on an 8 byte boundary.

Fields are placed in memory so that their offset is a multiple of their alignment.

## 2. Padding

If a field is not aligned, the compiler adds "empty" bytes (padding) before it: for example, 7 bytes of padding may be added after `int8` so that the next field of type `int64` starts at an address multiple of 8.

## 3. Structure size

The size of a structure is rounded up to a multiple of the maximum alignment of its fields. This ensures that arrays of structures are properly aligned.

## Example

```go
type ExampleStruct struct {
    Field1 int8
    Field2 int64
    Field3 int32
}
```

Example structure size: 24 bytes (1 byte + 7 paddings + 8 bytes + 4 bytes + 4 paddings).

## Why is this necessary?

Alignment improves performance by making the processor work more efficiently with memory. It minimizes the number of read/write operations and reduces errors related to unaligned data accesses.

# How to use

## Run project

You can launch the project from the `cmd/basics/structures/alignment` folder by the command:

```cmd
go run main.go
```

## Use entities

To use it in another project, you must do the following:

1. Import dependencies:

```go
import (
	"github.com/Mr-Filatik/go-knowledge-base/internal/basics/structures/alignment/printer"
	"github.com/Mr-Filatik/go-knowledge-base/internal/basics/structures/alignment/scanner"
)
```

2. Use in code:

```go
scan := scanner.ScanStructInfo(structInstanse, "Scan Name")
printer.PrintStructInfo(scan)
```

# Links

* [Struct Alignment and Padding in Go: A Comprehensive Guide](https://medium.com/@praffulmishra/struct-alignment-and-padding-in-go-a-comprehensive-guide-ae928d5a9d5e "Article on Medium")
* [Optimizing structures in Golang for efficient memory allocation](https://medium.com/nuances-of-programming/%D0%BE%D0%BF%D1%82%D0%B8%D0%BC%D0%B8%D0%B7%D0%B0%D1%86%D0%B8%D1%8F-%D0%B2%D1%8B%D1%80%D0%B0%D0%B2%D0%BD%D0%B8%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5-%D1%81%D1%82%D1%80%D1%83%D0%BA%D1%82%D1%83%D1%80-%D0%B2-golang-%D0%B4%D0%BB%D1%8F-%D1%8D%D1%84%D1%84%D0%B5%D0%BA%D1%82%D0%B8%D0%B2%D0%BD%D0%BE%D0%B3%D0%BE-%D1%80%D0%B0%D1%81%D0%BF%D1%80%D0%B5%D0%B4%D0%B5%D0%BB%D0%B5%D0%BD%D0%B8%D1%8F-%D0%BF%D0%B0%D0%BC%D1%8F%D1%82%D0%B8-92722244fd05 "Article on Medium")
* [Variable order in the Go golang structure](https://ru.stackoverflow.com/questions/1541584/%D0%9F%D0%BE%D1%80%D1%8F%D0%B4%D0%BE%D0%BA-%D0%BF%D0%B5%D1%80%D0%B5%D0%BC%D0%BD%D0%BD%D1%8B%D1%85-%D0%B2-%D1%81%D1%82%D1%80%D1%83%D0%BA%D1%82%D1%83%D1%80%D0%B5-go-golang "Question on StackOverflow")
