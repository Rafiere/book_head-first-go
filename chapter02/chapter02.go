package chapter02

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func Conditionals() {

	fmt.Println("\n--- CAPÍTULO 02 ---")

	/* Exibindo a data atual: */
	fmt.Println("\nExibindo a data atual:")

	currentTime := time.Now()
	fmt.Println(currentTime.Date())

	/* Comentários: */
	fmt.Println("\nComentários:")

	// Esse é um comentário de uma linha.
	/*
	   Esse é um
	   comentário de
	   várias linhas
	*/
}

func PassFail() {

	/* Verificador de Notas: */
	fmt.Println("\nVerificador de Notas:")

	fmt.Println("Insira a sua nota: ")

	/* Criando um novo buffer que permitirá lermos os caracteres digitados pelo teclado. */
	reader := bufio.NewReader(os.Stdin)

	/* Retornando tudo o que o usuário digitou até que ele aperte a tecla "Enter. */
	input, err := reader.ReadString('\n')

	/* Tratando possíveis erros: */
	if err != nil {
		log.Fatal(err)
	}

	/* Printando o que o usuário digitou: */

	fmt.Println("\nA nota do usuário é:", input)

	/* Retirando o "\n" da string: */
	input = strings.TrimSpace(input)

	/* Convertendo a string para um float64: */
	grade, err2 := strconv.ParseFloat(input, 64)
	if err2 != nil {
		log.Fatal(err2)
	}

	/* Verificando se o usuário passou: */
	var hasPassed bool

	if grade >= 60 {
		hasPassed = true
	} else {
		hasPassed = false
	}

	fmt.Println("Ele passou:", hasPassed)
}

func RandomNumberGame() {

	/* Advinhe o Número Aleatório: */
	fmt.Println("\nAdivinhe o Número Aleatório:")

	/* Gerando a seed para evitar números aleatórios: */
	seconds := time.Now().Unix()
	rand.Seed(seconds)

	targetNumber := rand.Intn(100) + 1

	for tries := 1; tries <= 10; tries++ {

		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Escolha um número: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guessIntInput, err2 := strconv.Atoi(input)
		if err2 != nil {
			log.Fatal(err2)
		}

		fmt.Println("Você tem", 10-tries, "tentativas!")

		if guessIntInput < targetNumber {
			fmt.Println("O seu número foi muito baixo!")
		} else if guessIntInput > targetNumber {
			fmt.Println("O seu número foi muito alto!")
		} else {
			fmt.Println("Você ganhou!")
			break
		}

		if tries == 10 {
			fmt.Println("Você perdeu. O número era ", targetNumber)
			break
		}
	}
}
