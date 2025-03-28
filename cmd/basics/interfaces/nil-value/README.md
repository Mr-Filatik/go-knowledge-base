# Interfaces and nil in Go

Inside Go, an interface variable is represented as a pair (type, value), where value is a specific value and type is the type of this value (in reality, it is a bit more complicated, but let's not dive into the maze).

As long as no value is assigned to the interface variable, both type and value are equal to nil, so the variable itself is considered equal to nil.

But as soon as a value is assigned to the interface variable, type ceases to be nil. Therefore, the variable is no longer equal to nil, even if value is equal to nil.

Checking for nil for interfaces requires special care. A simple `interfaceVar == nil` comparison may not be informative enough for interfaces containing nil values. A distinction must be made between a nil interface and an interface with a nil value.

As a solution, I propose two options:

1. Perform a nil value check before assigning a value to an interface variable;
2. Use `reflect.ValueOf(v).IsNil()`, but remember that using reflect incurs additional overhead.

# How to use

## Run project

You can launch the project from the `cmd/basics/interfaces/nil-value` folder by the command:

```cmd
go run main.go
```

# Links

* [Interfaces and nil in Go](https://antonz.ru/nil-interface/ "Article on Antonz")
* [Anatomy of interfaces in Go](https://habr.com/ru/articles/691956/ "Article on Habr")
