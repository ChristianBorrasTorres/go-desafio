package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aldogayaladh/go-desafio-test/internal/productos"
)

const (
	filename = "./productos.csv"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	var entrada string

	fmt.Print("Ingrese producto a buscar: ")

	_, err := fmt.Scanln(&entrada)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	storage := productos.Storage{
		Productos: readFile(filename),
	}

	// Crear canales para comunicarnos con las go rutinas
	canalProducto := make(chan productos.Producto)
	defer close(canalProducto)
	canalErr := make(chan error)
	defer close(canalErr)

	go func(chan productos.Producto, chan error) {

		producto, err := storage.BuscarProductoPorNombre(entrada)
		if err != nil {
			canalErr <- err
			return
		}

		canalProducto <- producto

	}(canalProducto, canalErr)

	select {
	case pr := <-canalProducto:
		fmt.Println(pr)
	case err := <-canalErr:
		log.Fatal(err)
		os.Exit(1)
	}

}

// readFile es una funcion que lee un archivo
func readFile(filename string) []productos.Producto {
	file, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	data := strings.Split(string(file), "\n")

	var resultado []productos.Producto
	for i := 0; i < len(data); i++ {

		if len(data[i]) > 0 {
			file := strings.Split(string(data[i]), ",")
			producto := productos.Producto{
				Nombre:      file[0],
				Descripcion: file[1],
				Cantidad:    file[2],
				Precio:      file[3],
			}
			resultado = append(resultado, producto)
		}

	}

	return resultado

}
