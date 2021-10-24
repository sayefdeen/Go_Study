# Go_Study

[Video](https://www.youtube.com/watch?v=YS4e4q9oBaU&ab_channel=freeCodeCamp.org)

Go is focusing in building servers and web applications, not client applications

## Why using GO?

1. Strong and statically language : It means the type of the variables can't change over time, all these variables have to be defined at compile time

2. Excellent community

3. Key features
   - Simplicity
   - Fast compile times
   - Garbage colleted
   - Build-in concurrency
   - Compile to standalone binaries

## Useful Resources

1. [Go Lang WebSite](https://golang.org/)

2. [Effective Go](https://golang.org/doc/effective_go)

## Go Commands

1. go run path to .go file : running the program

2. go build path to .go file : Building the program

3. go install : installing any outside packages

```Go
package main

import "fmt"

func main() {
	fmt.Println("Hello Go!")
}
```

## Variables

The compiler will figure out what is the type of the variables, each variable should be used, if it wasn't a compiler error will be thrown

```Go
var name type
var i int =42
// Or using this
i := 42
```

- Lower case variable it a package level.

- Upper case variable it will be exposed to anywhere outside the package. globally visible

- block scope, variables inside functions

|  int  |        Description        |
| :---: | :-----------------------: |
| int8  |        -128 to 127        |
| int16 |      -32768 to 32767      |
| int32 | -2147483648 to 2147483647 |
| int64 |         Too much          |

<p style='color:red'>You can't add two int with different types, you have to cast one type to another</p>

## Data Types

- string and bytes

Changing the string to bytes, and return it to string again

you can use that to send a string or file as a bytes into response

```Go
func method() {
	s := "This is a string"
	b := []byte(s)
	d := string(b)
	fmt.Println(b)
	fmt.Println(d)
}
```

Define a variable at compile time

```Go
func compileTime() {
	//  This will works since a will be created at compile time, the compiler will see when it is used and will define its type according to that
	const a = 42
	var b int16 = 27
	fmt.Println(a + b)
}
```

## Collections Types
