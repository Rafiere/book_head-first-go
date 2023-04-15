package chapter03

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

func FormattingValues() {

	fmt.Println("\n--- CAPÍTULO 03 ---")

	/* Formatando valores: */
	fmt.Println("\nFormatando valores")

	fmt.Println("Quase um terço: ", 1.0/3.0)
	fmt.Printf("Valor formatado: %0.2f\n", 1.0/3.0)
}

func SayHi() {
	fmt.Println("Hi!")
}

func PrintAPositiveNumber(number int) {

	/* Printando valores positivos: */
	fmt.Println("\nPrintando valores positivos")

	if number <= 0 {
		err := errors.New("O número deve ser maior que 0.")
		log.Fatal(err)
	}

	fmt.Println(number)
}

func Pointers() {

	/* Ponteiros: */
	fmt.Println("\nPonteiros")

	numero := 1
	fmt.Println(numero, "-", reflect.TypeOf(numero))
	fmt.Println(&numero, "-", reflect.TypeOf(&numero))
	fmt.Println("")

	/* Armazenando ponteiros em variáveis e obtendo os seus valores: */

	numero2 := 2
	numero2Ponteiro := &numero2

	fmt.Println(numero2, "-", reflect.TypeOf(numero2Ponteiro))
	fmt.Println("Valor em que o ponteiro está apontando: ", *numero2Ponteiro)
}
