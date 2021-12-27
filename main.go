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
	fmt.Println(a % b)

}