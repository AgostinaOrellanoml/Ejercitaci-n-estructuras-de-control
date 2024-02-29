package main

import "fmt"

func main() {
	var palabra string
	var contador int = 0

	fmt.Println("Ingrese una palabra: ")
	fmt.Scanf("%s", &palabra)
	fmt.Println("La palabra ingresada es: ", palabra)

	for i := 0; i < len(palabra); i++ {
		contador++

	}
	fmt.Println("La palabra tiene: ", contador, "letras")

}
