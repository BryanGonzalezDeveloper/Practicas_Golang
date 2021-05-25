package main

import (
	"fmt"
	"bufio"
	"os"
)

/*En Golang no existen bucle como while, do while.... Solo se utiliza el bucle for sin embargo puede ser
utilizado de maneras similares para emular esos comportamientos en golang el ciclo for es bastante flexible
*/


func main(){
	for_estilo_while()
}


func for_Normal(){
	for i := 0; i < 10; i++ {
		fmt.Println("Valor de i: ",i)
	}
}

func for_estilo_while()  {
	for{
		fmt.Println("Emulando while.")
		scanner:=bufio.NewScanner(os.Stdin)
		print("Escriba 'salir' para terminar el ciclo ")
		if  scanner.Scan() {
			if scanner.Text()=="salir"{
				break
			}
		}
	}
}