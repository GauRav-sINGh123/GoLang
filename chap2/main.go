package main

import "fmt"

func main() {
	var a int      // Declare without initialization
	var b int = 10 // Declare and initialize
	var c = 20     // Type inferred

	const pi = 3.14 // Constant cannot be changed
	pi = 4
	d := 30 // Short declaration. Type inferred as int

	var x, y, z int = 1, 2, 3  // Same type
	e, f, g := 4, "hello", 5.5 // Different types

	fmt.Println(a, b, c, d, pi, x, y, z, e, f, g)
}

// Block scope variable declaration
// func main() {
//     var (
//         name  string = "Gaurav"
//         age   int    = 25
//         score float64 = 99.5
//     )
//     fmt.Println(name, age, score)
// }
