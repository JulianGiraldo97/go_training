package main

import (
	"fmt"
	"hello_world/variables"
	"runtime"
)

func main() {
	estado, texto := variables.ConviertoaTexto(1558)
	fmt.Print(estado)
	fmt.Print(texto)
	if os := runtime.GOOS; os == "windows" || os == "OS X." {
		fmt.Println("Esto no es windows, es", os)
	} else {
		fmt.Println("Esto es windows, es", os)
	}
}
