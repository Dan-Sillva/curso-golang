package main

import (
	"modulo-test/auxiliar"
	"fmt"
	"github.com/badoux/checkmail"
)

func main () {
	auxiliar.Escrever()
	fmt.Println("vindo do main")

	erro0 := checkmail.ValidateFormat("devmail@mail.com")
	fmt.Println(erro0)

	erro1 := checkmail.ValidateFormat("1234")
	fmt.Println(erro1)
}