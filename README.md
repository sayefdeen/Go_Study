# Go_Study

[Video](https://www.youtube.com/watch?v=YS4e4q9oBaU&ab_channel=freeCodeCamp.org)

Go is focusing in building servers and web applications, not client applications

## GO?

1. Strong and statically language : It means the type of the variables can't change over time, all these variables have to be defined at compile time

2. Excellent community

3. Key features

   - Simplicity
   - Fast compile times
   - Garbage colleted
   - Build-in concurrency
   - Compile to standalone binaries

4. Go have Pointers such as C and C#

5. Go Have the concept on closures

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

If you don't have the data at declaring stage, you can use the build in `make()` function

```Go
func usingMake() {
	statePopulation := make(map[string]int)
	a := make([]int, 3)
}
```

- Arrays

This syntax represent creating an array with the same number of elements as in the initializing

```Go
func main() {
	grades := [...]int{97, 85, 93}
	fmt.Println("Grades ", grades)
	fmt.Println("Grades ", len(grades))
}
```

This syntax represent creating an array with specific number of elements

```Go
func main() {
	grades := [4]int
	fmt.Println("Grades ", grades)
```

This syntax represent coping an array, shallow copy not as a reference

a will hold the same value, and it will not be effected by the change in b

b will hold a copy from a

to override this behavior, we should use this symbol (&) before the assigning, its called **_address off operator_**

c will point to a in the memory, hench changing in c will affect a

```Go
func main() {
	a := [...]int{1, 2, 3}
	b := a
	c := &a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)
}
```

- Slice

Having an array with fixed size at compile time, make it limited for usage that's why we are going to use Slice

**Slice called reference type data**

```Go
func usingSlice() {
	a := []int{1,2,3,4,5,6,7,8,9,10}
	b := a[:] // Slice of all elements, same as the above array
	b := a[3:] // Slice from the 4th element to end, [4,5,6,7,8,9,10]
	b := a[:6] // Slice first 6 elements, [1,2,3,4,5,6]
	b := a[3:6] // Slice the 4th, 5th and 6th element, [4,5,6]

	// using spread operator to append new slice to this one

	a = a.append(a, []int{11,12,13}...)
}
```

- Maps

Its like an object, key, value Pair

Maps are reference type, such as Slice

```Go
func main() {
	statePopulate := map[string]int{
		"Jordan": 12345678,
		"Egypt":  87654321,
		"Syria":  34567789,
	}
	// This will print 0, since Jorda is not available in the map

	fmt.Println(statePopulate["Jorda"])

	// to prevent this to cause some confusion, we can use comma ok syntax
	// , ok
	// This will print 0 false, and we can check on the ok variable

	prop, ok := statePopulate["Jorda"]
	fmt.Println(prop, ok)
}
```

- Struct

Struct are value types, you can use the address operator again here

Its a data type can hold multiple data types, not a good idea to be used from maintenance view, if you change the order or adding a new data type, you should fix this on all the other file that use this struct

```Go
type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "Saif",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
		},
	}

	fmt.Println(aDoctor)
}
```

You can use it as an anonymous struct, only @runtime, when this function is used the struct will be created and removed from memory when the job is done

```Go
func main() {
	bDoctor := struct{ name string }{name: "Saif"}
	fmt.Println(bDoctor)
}
```

Struct doesn't support inheritance, but it support Composition

```Go
type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func something() {
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = true

	c := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Aus"},
		SpeedKPH: 47,
		CanFly:   false,
	}
}
```

## Conditions

- If Statement

General Syntax

```Go
/**
	if decleration; (condition) {
		whatever
	}
	*/
```

Example

```Go
func main() {
	statePopulate := map[string]int{
		"Jordan": 12345678,
		"Egypt":  87654321,
		"Syria":  34567789,
	}

	if pop, ok := statePopulate["Jordan"]; ok {
		fmt.Println(pop)
	}
}
```

- Switch Statement

```Go
func switchStm() int {
	switch 3 {
	case 1, 2, 4:
		return 1

	case 3:
		return 3

	}
	return 0
}
```

## Loops

Same as every loop ever, but we have something called Label

This code will break the outer loop

```Go
func main() {
Loop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}
}
```

For Range Loop

```Go
array := []int{1, 2, 3}
	for index, value := range array {
		fmt.Println(index, value)
	}

	array := []int{1, 2, 3}
	for _, value := range array {
		fmt.Println( value)
	}
```

## KeyWords

- Defer

**Invoke** a function, but delay its execution to a specific time, the defer functions follow the Queue behavior

Since the function is invoked, it will hold the values on time of invoking, not the time of execution, this code will print `Start`

```Go
func main() {
	a := "Strart"
	defer fmt.Println(a)
	a = "End"
}

```

Most used when you open a resource that needs to be closed such as reading file stream

defer will move the block just before the function return, at it will execute according to LIFO

```Go
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", robots)
}
```

- Panic 🙀

All panic statements happened after the defer statements executed

- Recover

The code below will be executed in this order

- Start will print

- about the panic will print

- defer will delay the execution of anonymous function

- panicker will panic, throwing an error Something bad happened

- the defer function will executed

- recover will return an error

- the condition will be true

- an error will be printed

- End will be printed from main

```Go
import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Start")
	panicker()
	fmt.Println("End")
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)
			panic(err)
		}
	}()
	panic("Something bad happened")
	fmt.Println("done Panicking")
}
```

In case We can't handle the error we can panic again 🙀

## Pointers

As we said, in go we use pointers to point to the variable reference such as C and C++

We can't apply pointer arithmetic, which is applying math operators in the references values

Slice and Maps for example, are by default using pointers in shadow

- & reference operator

```Go
func main() {
	a := []int{1, 2, 3}
	b := a
	fmt.Println(a, b)
	a[1] = 42
	fmt.Println(a, b)
	greeting:= "Hello"
	name := "Stacey"
	sayGreeting(&greeting,&name)
}

func Maps() {
	a := map[string]string{"foo": "bar", "baz": "buz"}
	b := a
	fmt.Println(a, b)
	a["foo"] = "qux"
	fmt.Println(a, b)
}

- * called a reference destruct, without it the reference representation will be printed instead of the value itself

func sayGreeting(greeting *string, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}

```

## InterFaces

Interfaces describe behaviors, store method definition

```Go
func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello World"))
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct {
}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
```

- Use many, Small Interfaces

- Don't export interfaces fro types that will be consumed

- Design functions and methods to receive interfaces whenever possible

## Goroutines

go keyword before the function, this will tell go to start a [greenThread](https://en.wikipedia.org/wiki/Green_threads) and run this function if this thread

- Sync

We can use the [sync](https://pkg.go.dev/sync@go1.17.2) package from go to handle all the goroutines to be synced

## Channels
