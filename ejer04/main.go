package main

import (
	"fmt"
	"strconv"
	"bufio"
	"os"
)

var numero1 int
var numero2 int
var texto string
func main() {

	fmt.Println("Ingrese un numero entero: ")
	fmt.Scanln(&numero1)
	fmt.Println("Ingrese un segundo numero entero: ")
	fmt.Scanln(&numero2)
	println("Numero 1: "+strconv.Itoa(numero1))
	println("Numero 2: "+strconv.Itoa(numero2))
	println("Suma: "+strconv.Itoa(numero1+numero2))

	println("Descripcion: ")
	scanner:=bufio.NewScanner(os.Stdin)
	if scanner.Scan(){
		texto=scanner.Text()
	}
	println("Texto ingresado= "+texto)
}
