package main
import ("fmt")


//declaração de variáveis do tipo inteiro
var numero1 int = 1;
var numero2 int = 2;
var resultado int = 0;

// declaração de constante
const PI float32 = 3.14;

func main() {
	var a string = "Lucas";
	var b string = "Pedro";
	var c string = "Maria";

fmt.Println("Olá", a, b, c, PI); // aqui main imprime as variáveis e a constante

  soma() // aqui main chama a função soma abaixo
}

//declaração da função soma
func soma() {
	resultado = numero1 + numero2;
	fmt.Println("O resultado da soma é:", resultado);
}