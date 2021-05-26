package main
import(
	"fmt"
	
	)

func main(){
	fmt.Println("Ingrese un valor numerico")
	var num float32
	fmt.Scanln(&num)
	resultado:=iva(num)
	println("El iva de %d es igual a %d",num,resultado)
}

func iva(numero float32) float32{
	fmt.Println("Esta devolvera el iva (16%) de: ",numero)
	return numero*0.16
}