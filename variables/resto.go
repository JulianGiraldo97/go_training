package variables

import (
	"fmt"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time

func RestoVariables() {
	Nombre = "Pedro"
	Estado = true
	Sueldo = 0
	Fecha = time.Now()
	fmt.Println(Nombre, Estado, Sueldo, Fecha)
}
