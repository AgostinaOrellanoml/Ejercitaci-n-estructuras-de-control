package main

import "fmt"

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

func main() {
	fmt.Println("La edad de Benjamin es: ")
	fmt.Println(employees["Benjamin"]) //Imprimo la edad de un empleado
	/*if employees > 21 {  VER COMO IMPRIMO LOS MAYORES

	}
	*/
	//Elimino a un empleado de la lista
	employees["Pedro"] = 30
	delete(employees, "Pedro")
	fmt.Println((employees))
	//Agrego un elemento
	employees["Federico"] = 25
	fmt.Println(employees)
}
