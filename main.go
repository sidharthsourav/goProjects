package main

import (
	"fmt"
	"goProjets/helper"
	"goProjets/stringutils"
)

func main() {
	fmt.Println("Hello World")
	message := helper.PrintHelper("mainCode printMessage")
	fmt.Println(message)
	helper.Greeting(message)
	fmt.Println(stringutils.Reverse(message))
	if secret, key := helper.MultipleReturner(); key == 13 {
		fmt.Println(secret)
	}

	fmt.Println("From-StringReverse-fn: " + stringutils.StringReverse(message))
	fmt.Println(helper.BasicArithmetic(9, 3))
	fmt.Println(helper.C2FConvertor(37.1))
	fmt.Println(helper.CircleAreaCalculator(67.1))
	age := 40
	name := "ACG"

	{
		age, newName := 40, "johnDoe"
		fmt.Println(age, newName)
	}
	fmt.Println(name, age)
	name = "Raj"
	name, age2 := "kumar", 20
	fmt.Println(name, age2)
}
