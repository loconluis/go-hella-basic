package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Declaracion de variables
	var x int
	x = 8
	y := 7

	fmt.Println(x)
	fmt.Println(y)

	myValue, err := strconv.ParseInt("L", 0, 64)
	// Formas de manejar errores, de manera EXPLICITA
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("myValue: ", myValue)
	}

	// Maps Estructuras de llave valor
	//           map [llave][tipos de valor que obtendra la llave]
	// m := make(map[string]int)
	m := make(map[string]int)
	m["key"] = 6

	fmt.Println("M[KEY]", m["key"])

	// Slices
	s := []int{1, 2, 3}

	for index, value := range s {
		fmt.Println("Index", index)
		fmt.Println("Value", value)
	}
	fmt.Println("************************")
	// Agregar valores a un slice
	s = append(s, 16)
	for index, value := range s {
		fmt.Println("Index", index)
		fmt.Println("Value", value)
	}
}
