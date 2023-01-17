package main

import (
	"fmt"
)

type user struct {
	nome 		string
	idade		uint8
}

func main() {
	fmt.Println("Structs")

	var us user
	us.nome = "Felipe Neto"
	us.idade = 5
	fmt.Println(us)

	us2 := user{"Jonas Neto", 12}
	fmt.Println(us2)

	us3 := user{nome: "Guts Neto"}
	fmt.Println(us3)
}