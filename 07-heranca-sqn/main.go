package main

import (
	"fmt"
)

type pessoa struct {
	nome		string
	idade		uint8
}

type ti_man struct {
	pessoa
	email		string
	cargo		string
	salario		float32
}

func main() {
	// var user ti_man
	
	pessoa := pessoa{"Ricardo", 12}
	fmt.Println(pessoa)

	user := ti_man{pessoa, "rick@sancherz.com", "GO SENIOR ENGENNIER", 9000.41}
	fmt.Println(user)
	fmt.Println(user.nome)
}