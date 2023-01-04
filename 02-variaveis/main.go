package main

import (
	"fmt"
)

func main()  {
	var str string = "variavel em forma de string"
	str2 := "variavel em forma de str"

	fmt.Println(str+"\n"+str2)
	//-----------------------------------------------

	var (
		var1 string = "nova string"
		var2 int = 12
	)

	fmt.Println("\n",var1,"\n",var2)
	//-----------------------------------------------

	g1, g2, g3 := "g1", 2, "g3"

	fmt.Println("\n",g1, g2, g3)
	//-----------------------------------------------

	const constante string = "constante"
	
	fmt.Println(constante)
	//-----------------------------------------------

	str, str2 = str2, str // inversao de valores de variaveis
	fmt.Println(str+"\n"+str2)
}