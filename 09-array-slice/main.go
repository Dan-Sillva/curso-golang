package main

import (
	"fmt"
)

func main() {
	var array1 [5]string
	array1[0] = "p1"
	fmt.Println(array1, len(array1))

	array2 := [2]string{"p1", "p2"}
	fmt.Println(array2, len(array2))

	array3 := [...]string{"p1", "p2", "p3", "p4", "p5", "p6"}
	fmt.Println(array3, len(array3))

	slice := []int{1, 3, 7, 12, 5}
	fmt.Println(slice, len(slice))

	slice = append(slice, 18) 
	fmt.Println(slice, len(slice))

	slice2 := array3[1:3]
	fmt.Println(slice2, len(slice2))
}