package main

import (
	"fmt"
	"hello_world/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(1558)
	fmt.Print(estado)
	fmt.Print(texto)
}
