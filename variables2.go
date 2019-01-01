package main

import "fmt"

func main() {

	var a = "initial variable" // Single variable declaration
	fmt.Println(a)

	var b, c int = 1, 2 // Multi variable declaration in single line
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int // Without  initialization variables become zero-valued by default
	fmt.Println(e)

	f := "short type declaration" // Shorter way of declaring and initializing the variable
	fmt.Println(f)

}

//  Questions
// How to declare a float variable ?
// What are the other types of variables available ?
// How is shorter declaration different from normal declaration?
// Why is int needed to be specified and not string ?
// What happens if we define by "var e string" ?
