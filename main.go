package main

import "fmt"

var l int = 42 //must declare in this way outside of func
var (
	actorName string = "elisabeth sladen"
	companion string = "some person"
	doctorNumber int = 3
	season int = 11
) //can wrap in var blocks like this for less clutter

func main() {
	fmt.Println(" hello world")

	var i int
	i = 42
	fmt.Println(i)

	// i := 42 will let go compiler recognize what type this is

	var j int = 27
	k := 99
	fmt.Println(i,j,k)
	fmt.Printf("%v, %T", j , j) //print var j and type of j
	fmt.Printf("%v, %T", l, l)

}