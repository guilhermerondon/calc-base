package main

import (
	"fmt"
)

func main() {

	var num1, num2 float64
	var operacao string

	fmt.Print("Digite o primeiro numero: ")
	fmt.Scan(&num1)

	fmt.Print("Digite o segundo numero: ")
	fmt.Scan(&num2)

	fmt.Print("Digite a operação (+, -, *, ;): ")
	fmt.Scan(&operacao)

	switch operacao {
	case "+":
		fmt.Printf("Resultado: %.2f\n", num1+num2)
	case "-":
		fmt.Printf("Resultado: %.2f\n", num1-num2)
	case "*":
		fmt.Printf("Resultado: %.2f\n", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("REsultado: %.2f\n", num1/num2)
		} else {
			fmt.Printf("erro; div por zero!")
		}

	default:
		fmt.Println("Operacao invalida!")

	}
}
