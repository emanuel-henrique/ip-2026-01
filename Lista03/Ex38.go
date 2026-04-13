package main

import (
	"fmt"
	"strconv"
)

func main() {
	var cpf string

	fmt.Print("Digite um CPF: ")
	fmt.Scan(&cpf)

	isValidCPF := validarCPF(cpf)

	if isValidCPF {
		fmt.Println("CPF Válido!")
	} else {
		fmt.Println("CPF Inválido!")
	}
}

func validarCPF(cpf string) bool {
	var veficador1, veficador2 int
	count := 10
	soma := 0

	for i := 0; i <= len(cpf)-3; i++ {
		num, _ := strconv.Atoi(string(cpf[i]))
		num *= count
		soma += num
		count--
	}

	resto := soma % 11

	if resto < 2 {
		veficador1 = 0
	} else if resto >= 2 {
		veficador1 = 11 - resto
	}

	count = 11
	soma = 0

	for j := 0; j <= len(cpf)-3; j++ {
		num, _ := strconv.Atoi(string(cpf[j]))
		num *= count

		soma += num
		count--
	}

	soma += (veficador1 * 2)
	resto = soma % 11

	if resto < 2 {
		veficador2 = 0
	} else if resto >= 2 {
		veficador2 = 11 - resto
	}

	cpfVerificardor1, _ := strconv.Atoi(string(cpf[9]))
	cpfVerificardor2, _ := strconv.Atoi(string(cpf[10]))

	if cpfVerificardor1 == veficador1 && cpfVerificardor2 == veficador2 {
		return true
	} else {
		return false
	}
}
