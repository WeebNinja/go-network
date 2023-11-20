package main

import (
	"fmt"
)

// list of the data types
// bool
// string
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte // alias for uint8
// rune // alias for int32

// float32 float64
// complex64 complex128
// var c complex64 = 5+5i

// main is the entry point for the application.

func main() {
	var t bool = true
	var f bool

	fmt.Println("t is ", t)
	fmt.Println("f is ", f)

	var age int = 40
	var favNum float64 = 1.6180339

	var complex complex128 = 5 + 5i

	// What is rune?
	//	- rune is an alias for int32
	//	- represents a Unicode code point
	//	- when you iterate over a string in Go, you get back runes
	var r rune = 10 // rune is an alias for int32

	fmt.Println("age is ", age, "favNum is ", favNum)
	fmt.Println("complex128 is ", complex)
	fmt.Println("rune is ", r)

	var myName string = "อนุสรณ์ ใจแก้ว"
	fmt.Println(myName + " is a robot")
	fmt.Println(myName[3])
	fmt.Println("length of myName is ", len(myName))

	// Arrays variables
	var arry5 [5]float64 // array of 5 float64
	arry5[0] = 98.7      // assign value to array
	arry5[1] = 93.2      // assign value to array
	arry5[2] = 77.2      // assign value to array
	arry5[3] = 83.7      // assign value to array
	arry5[4] = 88.2      // assign value to array
	fmt.Println(arry5)
	fmt.Println("length of arry5 is ", len(arry5))
	fmt.Println("arry5[3] is ", arry5[3])

	// another way to declare array
	arry3 := [3]float64{98, 93, 77}
	fmt.Println(arry3)

	// slice is a segment of an array
	// slice is a reference type
	// slice is a pointer to an array
	// slice is a length and a capacity
	// slice is a dynamic size
	// slice is a flexible view into the elements of an array

	var s []float64 = arry5[0:2] // slice of an array
	fmt.Println(s)

	// struct is a collection of fields
	// struct is a user-defined type
	// struct is a composite type
	// struct is a collection of fields
	// struct is a blueprint

	type Person struct {
		name string
		age  int
	}

	var p Person
	p.name = "Anusorn"
	p.age = 40
	fmt.Println(p)

	// pointer is a variable that stores the memory address of another variable
	// pointer is a reference type
	// pointer is a value that points to the address of a value

	var x int = 5
	var xPtr *int = &x
	fmt.Println(xPtr) // print the memory address of x

	if 20 > 18 {
		fmt.Println("20 is greater than 18")
	}

	temperature := 14
	if temperature > 15 {
		fmt.Println("It is warm out there")
	} else {
		fmt.Println("It is cold out there")
	}

	time := 22
	if time < 10 {
		fmt.Println("Good morning.")
	} else if time < 20 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}

	num := 20
	if num >= 10 {
		fmt.Println("Num is more than 10.")
		if num > 15 {
			fmt.Println("Num is also more than 15.")
		}
	} else {
		fmt.Println("Num is less than 10.")
	}

	day := 4

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	}

	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

}
