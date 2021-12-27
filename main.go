package main

import "fmt"

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

}