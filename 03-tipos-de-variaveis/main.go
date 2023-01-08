package main

import (
	"fmt"
	"errors"
)

func main (){
	var number0 int8 = 123
	var number1 int16 = 123
	var number2 int32 = 123
	var number3 int64 = 123
	fmt.Println(number0, number1, number2, number3)

	var num uint = 10 // numeros sem sinais de + e -
	fmt.Println(num)
	
	// alias para o int32, usado mais com numeros que representam caracteres
	var num2 rune = 32
	fmt.Println(num2)

	// alias para uint8 (uint de 8 bits) que tem o nome de ?
	var num3 byte = 100
	fmt.Println(num3)

	// float32 ou float64 para pontos flutuantes
	var num4 float32 = 12.2
	var num5 float64 = 12331123.123
	fmt.Println(num4,num5)
	
	// string
	var str string = "frase qualquer"
	fmt.Println(str)

	// boolean
	var boll bool = true
	fmt.Println(boll)

	// error
	var erro error = errors.New("pra largar de ser trouxa")
	fmt.Println(erro)

}