package chapter01

import (
	"fmt"
	"reflect"
	"strings"
)

func HelloWorld() {

	fmt.Println("\n--- CAPÍTULO 01 ---")

	fmt.Println("\nHello World:")

	/* Common Hello World */
	fmt.Println("Hello, world!")

	/* Hello World - Title */
	fmt.Println(strings.Title("hello world"))

	fmt.Println("\nRunes:")

	/* Runas */
	fmt.Println('A') //Imprimirá 65, que é o ASCII.

	fmt.Println("\nNumbers:")

	/* Números */
	fmt.Println(45)

	/* Booleans */
	fmt.Println("\nBooleans:")

	fmt.Println("45 é maior que 25? - ", 45 > 25)
}

func Variables() {

	/* Variáveis */
	fmt.Println("\nVariables:")

	var quantity int
	var length, width float64
	var customerName string

	quantity = 2
	length, width = 1.52, 2.45
	customerName = "Rachel"

	fmt.Println("Quantidade:", quantity)
	fmt.Println("Length:", length)
	fmt.Println("Width:", width)
	fmt.Println("Customer Name:", customerName)

	/* Declarando e atribuindo valores em uma variável: */
	fmt.Println("\nDeclarando e atribuindo valores em uma variável:")
	quantity2 := 4
	length2, width2 := 1.52, 2.45
	customerName2 := "Rachel 2"

	fmt.Println("Quantidade 2:", quantity2)
	fmt.Println("Length 2:", length2)
	fmt.Println("Width 2:", width2)
	fmt.Println("Customer Name 2:", customerName2)

	/* Convertendo tipos: */
	fmt.Println("\nConvertendo tipos:")

	myIntegerVariable := 2
	fmt.Println(reflect.TypeOf(myIntegerVariable)) //int
	myFloatVariable := float64(myIntegerVariable)
	fmt.Println(reflect.TypeOf(myFloatVariable)) //float
}
