# Arrays

Basic rules for working with arrays:

- An array contains the values ​​of the array, not a reference to the first element of the array;
- The length of an array in Go is part of its type;
- If you do not specify explicit values for array elements, all elements are initialized with the default value for the type (for example: 0 for int, etc.);
- Array element indexing starts with the 0th element;
- Array element values can be changed by accessing them by index;
- Arrays are copied and assigned by value, not by reference;
- Arrays can only be compared if they have the same type and length.

# How to use

## Run project

You can launch the project from the `cmd/examples/arrays` folder by the command:

```cmd
go run main.go
```

# Links

* [About Arrays and Slices in Go](https://habr.com/ru/articles/739754/ "Article on Habr")

Press **[Back](../README.md "Return to previous README.md")** to go to the previous `README.md` file.