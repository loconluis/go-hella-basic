package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// Declaracion de variables
	fmt.Println("****VARIABLES****")
	var x int
	x = 8
	y := 7

	fmt.Println(x)
	fmt.Println(y)

	myValue, err := strconv.ParseInt("7", 0, 64)
	// Formas de manejar errores, de manera EXPLICITA
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("myValue: ", myValue)
	}

	// Maps Estructuras de llave valor
	//           map [llave][tipos de valor que obtendra la llave]
	// m := make(map[string]int)
	fmt.Println("****MAPS****")
	m := make(map[string]int)
	m["key"] = 6

	fmt.Println("M[KEY]", m["key"])

	// Slices
	fmt.Println("****SLICES****")
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
	fmt.Println("****GO ROUTINES****")
	c := make(chan int)

	// Go Routines
	go doSomething(c)
	<-c

	fmt.Println("****APUNTADORES****")
	g := 25
	fmt.Println("g: ", g)
	// APUNTADOR se hace con el simbolo & sobre el espacio en memoria
	// APUNTADOR se hace con el simbolo * devolver el valor de ese espacio
	h := &g
	fmt.Println("h: ", h)
	fmt.Println("h: ", *h)
}

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	// Enviara un valor a un canal
	c <- 1
}
