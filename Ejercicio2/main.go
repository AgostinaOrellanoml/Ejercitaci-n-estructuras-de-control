package main

import "fmt"

var edad int
var empleado bool = true
var antiguedad int
var sueldo float32

func main() {
	fmt.Println("Ingrese su edad: ")
	fmt.Scanf("%d", edad)

	if edad > 20 {
		fmt.Println("Ingrese si esta empleado: true,false ")
		fmt.Scanf("%s", empleado)
		//if empleado == True {   //ver como se pone bien usar un string
		fmt.Println("Ingrese años de antiguedad:")
		fmt.Scanf("%d", antiguedad)
		if antiguedad > 1 {
			fmt.Println("Ingrese el monto de su sueldo: ")
			fmt.Scanf("%f", sueldo)
		} else if sueldo > 1000 {
			fmt.Println("No se le cobra interes a su prestamo")
		} else {
			fmt.Println("Se le cobrara interes en su prestamo")
		}
	} else {
		fmt.Println("No le podemos otorgar su prestamo")
	}
}
