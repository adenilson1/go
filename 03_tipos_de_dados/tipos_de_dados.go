package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Numeros Inteiros")
	// numeros inteiros
	// int8, int16, int32, int64

	var numero1 int16 = 100
	var numero2 int64 = 10000000000000000
	var numero3 uint32 = 10000
	numero := 10000000000000000

	//alias
	// rune = 12345
	var numero4 rune = 12345	

	// BYTE = uint8
	var numero5 byte = 123

	fmt.Println(numero1)
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(numero5)
	fmt.Println(numero)

	fmt.Println("----------------------------")
	fmt.Println("Numeros Reais")
	// numeros reais
	//float32, float64

	var numeroreal1 float32 = 123.45
	fmt.Println(numeroreal1)

	var numeroreal2 float64 = 12300000000.45
	fmt.Println(numeroreal2)

	numeroreal3 := 12345.568
	fmt.Println(numeroreal3)

	fmt.Println("----------------------------")
	fmt.Println("Strings")

	// Strings => uso de aspas duplas
	// nao existe tipo char, no go considera alguma coisa
	// parecida como char com uso de aspas simples
	// char no go sao numeros na tabela asc, logo sao do tipo int
	
	var str1 string = "casa bonita"
	fmt.Println(str1)

	 str2 := "casa feia"
	fmt.Println(str2)

	char := 'c'
	fmt.Println(char)

	fmt.Println("----------------------------")
	fmt.Println("Valores Vazio, Zero, False e Error")

	// valor de uma variavel nao iniciada

	var texto string // vazio
	var num int16 // zero
	var boo1 bool // false
	// valor error
	var erro error // nil
	var erro1 error = errors.New("Error interno")
	fmt.Println(texto)
	fmt.Println(num)
	fmt.Println(boo1)
	fmt.Println(erro)
	fmt.Println(erro1)

	fmt.Println("----------------------------")
	fmt.Println("Booleanos")

	var booleano1 bool = true // ou false
	fmt.Println(booleano1)


}