package main

import "fmt"

var l int = 42 //must declare in this way outside of func
var (
	actorName string = "elisabeth sladen"
	companion string = "some person"
	doctorNumber int = 3
	season int = 11
) //can wrap in var blocks like this for less clutter and for organization

var I int = 42 //upper case, globally visable. //in main func -block scope, never visible outside block, lower case package, useable in block but not outside pkg

func main() {
	fmt.Println(" hello world")

	var i int
	i = 42
	fmt.Println(i)

	// i := 42 will let go compiler recognize what type this is

	var j int = 27
	k := 99
	fmt.Println(i,j,k)
	fmt.Printf("%v, %T\n", j , j) //print var j and type of j
	fmt.Printf("%v, %T\n", l, l)

	n := 1 == 1
	fmt.Printf("%v, %T/n", n, n) 

	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b) //go will error out with diff int types(int, uint8, 16, 32, 64) and other types like float. 

	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b)
	fmt.Println(a << 3)
	fmt.Println(a >> 3)

	m := 3.14
	m = 13.7e73
	m = 2.1E14
	var o float64 = 3.14
	fmt.Printf("%v, %T",m,m)
	fmt.Printf("%v, %T",o,o)

	var p complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", p , p)


}