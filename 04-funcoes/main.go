package main 

import (
	"fmt"
)

func soma (n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculos (n1, n2 int8) (int8, int8) { // funcao que retorna dois valores
	soma := n1 + n2
	subs := n1 - n2
	return soma, subs
}

func main (){
	fmt.Println(soma(2,3))

	var f = func () {
		fmt.Println("funcao f")
	}

	f()
	fmt.Println(calculos(5,5))
}