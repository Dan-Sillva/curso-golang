package main

import (
	"fmt"
)

func main() {
	// operacoes so sao feitas com variaveis de mesmo tipo
	// int16 + int32 = erro
	n, m := 7, 6 
	fmt.Println(n + m)
	fmt.Println(n - m)
	fmt.Println(n * m)
	fmt.Println(n / m)
	fmt.Println(n % m)

	//---------------------------------------------------

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	//---------------------------------------------------

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	//---------------------------------------------------

	numero := 10

	numero++
	numero--
	numero += 2
	numero -= 2
	numero *= 2
	numero /= 2

	fmt.Println(numero)
}