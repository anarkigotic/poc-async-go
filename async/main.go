package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // Espera que una colección de goroutines se ejecute

func saludar(nombres []string) {
	defer wg.Done() // Disminuye en uno la cantidad de goroutines a esperar
	for _, nombre := range nombres {
		fmt.Printf("Hola %s\n\r", nombre)
		time.Sleep(1 * time.Second)
	}
}

func despedir(nombres []string) {
	defer wg.Done() // Disminuye en uno la cantidad de goroutines a esperar
	for _, nombre := range nombres {
		fmt.Printf("Adiós %s\n\r", nombre)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	wg.Add(2) // Añade la cantidad de goroutines a esperar

	nombres := []string{"Orlando", "Daniela", "José",
		"Carlos", "Andrea", "David", "Carmen"}

	go saludar(nombres)
	go despedir(nombres)

	wg.Wait() // Espera hasta que se ejecuten las goroutines
	fmt.Println("Fin del programa")
}
