# Nil-recipient method in Go

The function can be called even on an empty pointer.

In this case, in order not to get `panic` it is necessary to add a check of the recipient for nil at the beginning of the function, if the method accesses fields of the structure.

This happens because methods in Go are bound to types rather than specific values. A method has a receiver, some variable that can be nil.

If a method is defined for a pointer type (e.g., `*MyStruct`), Go allows you to call this method even on a nil pointer. This is possible because methods in Go do not require the object itself to be initialized.

Instead, the method accesses the pointer (even if it is nil). Important: if the method does not access the structure fields, it can be executed successfully.

The ability to call methods on nil pointers can be useful, for example, to implement methods that check the state of an object.

# How to use

## Run project

You can launch the project from the `cmd/basics/structures/nil-pointer-func` folder by the command:

```cmd
go run main.go
```

# Links

* [Nil-recipient method in Go](https://antonz.ru/nil-method-receiver/ "Article on Antonz")
