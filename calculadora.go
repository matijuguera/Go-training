package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
temas aplicados
-Leer inputs de consola
-Printear en la consola resultados
-if y manejo de errores
-switch
-structs y receivers

Ejercicio
hacer una calculadora que +, -, /, * aplicando los conceptos aprendidos
*/

func main() {
	firstNumber := scanNumber("Ingrese el primer valor numerico: ")
	operatorValue := scanValue("Ingrese el operador: ")
	SecondNumber := scanNumber("Ingrese el segundo valor numerico: ")
	calc := calc{}
	calc.doOperation(operatorValue, firstNumber, SecondNumber)
}

type calc struct{}

func (calc) doOperation(operator string, number1 int, number2 int) {
	switch operator {
	case "+":
		fmt.Println(number1 + number2)
	case "-":
		fmt.Println(number1 - number2)
	case "*":
		fmt.Println(number1 * number2)
	case "/":
		fmt.Println(number1 / number2)
	default:
		fmt.Println(operator, " No esta soportado")
	}
}

func scanValue(label string) string {
	fmt.Println(label)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()
}

func scanNumber(label string) int {
	value := scanValue(label)
	number, _ := strconv.Atoi(value)

	return number

}
