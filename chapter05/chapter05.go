package chapter05

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func CreateArrays() {

	/* Creating Arrays: */
	fmt.Println("\nCreating Arrays")

	var notes [7]string

	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"

	fmt.Println(notes[0])
	fmt.Println(notes[1])

	/* Literal Arrays: */
	fmt.Println("\nLiteral Arrays")

	notes2 := [7]string{"do", "re", "mi", "fa", "sol", "la", "si"}

	fmt.Println(notes2)

	/* Formatando como o array aparece no código: */
	fmt.Println("\nFormatando como aparece no código")

	notes3 := [7]string{"do", "re", "mi", "fa", "sol", "la", "si"}

	fmt.Printf("\n%#v\n", notes3)

	/* Obtendo o tamanho do array: */
	fmt.Println("\nObtendo o tamanho do array: ")

	fmt.Printf("\nO array %#v tem %d elementos.\n", notes3, len(notes3))

	/* Obtendo os elementos do array (forma insegura - for): */
	fmt.Println("\nObtendo os elementos do array (forma insegura - for): ")

	for i := 0; i < len(notes3); i++ {
		fmt.Println(i, " - ", notes3[i])
	}

	/* Obtendo os elementos do array (forma segura - for-range): */
	fmt.Println("\nObtendo os elementos do array (forma segura - for-range): ")
	for index, value := range notes3 {
		fmt.Println(index, " - ", value)
	}
}

func calcAverage(array [3]float64, err error) {

	if err != nil {
		log.Fatal(err)
	}

	sum := 0.00

	for _, number := range array {
		sum += number
	}

	average := sum / float64(len(array))

	fmt.Println("A soma dos valores no arquivo é de ", sum)
	fmt.Println("A média dos valores no arquivo é de ", average)
}

func SumValuesInAFile() {

	/* Obtendo os valores em um arquivo: */
	fmt.Println("\nObtendo os valores em um arquivo: ")

	/* Nesse array, ficarão armazenados os números do arquivo. */
	var numbers [3]float64

	/* Estamos abrindo o arquivo. */
	file, err := os.Open("chapter05/data.txt")
	if err != nil {
		log.Fatal(err)
	}

	/* Essa variável armazenará o índice do array. */
	i := 0

	scanner := bufio.NewScanner(file)

	/* Estamos convertendo cada linha do arquivo, que é uma string, para um valor em float64. */
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}
		/* Estamos movendo o array para o próximo índice. */
		i++
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	if scanner.Err() != nil {
		log.Fatal(err)
	}

	/* Se nenhum erro ocorrer, estamos retornando o array de números e um erro com "nil". */

	calcAverage(numbers, nil)
}
