package main

import (
	"github.com/rafiere/head-first-go/chapter01"
	"github.com/rafiere/head-first-go/chapter02"
	"github.com/rafiere/head-first-go/chapter03"
	"github.com/rafiere/head-first-go/chapter05"
)

func main() {

	/* Chapter 01 */

	chapter01.HelloWorld()
	chapter01.Variables()

	/* Chapter 02 */

	chapter02.Conditionals()
	//chapter02.PassFail()
	//chapter02.RandomNumberGame()

	/* Chapter 03 */
	chapter03.FormattingValues()
	chapter03.SayHi()
	chapter03.PrintAPositiveNumber(10)
	chapter03.Pointers()

	/* Chapter 04 */

	/* Chapter 05 */

	chapter05.CreateArrays()
	chapter05.SumValuesInAFile()
	//fmt.Println(os.Getwd())

	/* Chapter 06 */
}
