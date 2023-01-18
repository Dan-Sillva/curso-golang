package main 

import (
	"fmt"
)

func main() {
	var n1 int = 10
	var n2 int = n1

	fmt.Println(n1, n2)
	
	n1++
	fmt.Println(n1, n2)

	// PONTEIROS
	var n3 int
	var p1 *int

	n3 = 100
	p1 = &n3
	
	fmt.Println(n3, p1)
	fmt.Println(n3, *p1)

	n3 += 50
	fmt.Println(n3, p1)
	fmt.Println(n3, *p1)
}