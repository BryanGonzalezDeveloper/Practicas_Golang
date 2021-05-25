package main
import "fmt"
//Declaracion de variable
var numero int
var texto string
var status bool

func main(){
var index int16
index=15
var contenedor int32
contenedor=int32(index)
fmt.Printf("Tipo de la variable index: %T\n",index)
fmt.Printf("Tipo de la variable contenedor: %T\n",contenedor)
pruebaTexto:="Texto de prueba"
fmt.Println(pruebaTexto)
metodoPrueba()
}

func metodoPrueba()  {
	fmt.Println("Funcionamiento de una funcion")
}