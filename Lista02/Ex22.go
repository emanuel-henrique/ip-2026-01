package main

import (
	"fmt"
)

func main() {
	const SalarioMinimo = 788.00
	const ValorHoraExtra = 10.00

	var matricula int
	var horasExtras float64
	var inss, ir float64

	fmt.Println("=== CÁLCULO DE FOLHA DE PAGAMENTO ===")

	fmt.Print("Matrícula do funcionário: ")
	fmt.Scan(&matricula)
	fmt.Print("Quantidade de horas-extras: ")
	fmt.Scan(&horasExtras)

	salarioHoraExtra := horasExtras * ValorHoraExtra

	salarioBruto := (3 * SalarioMinimo) + salarioHoraExtra

	if salarioBruto > 1500.00 {
		inss = salarioBruto * 0.12
	} else {
		inss = 0.0
	}

	if salarioBruto > 2000.00 {
		ir = salarioBruto * 0.20
	} else {
		ir = 0.0
	}

	salarioLiquido := salarioBruto - inss - ir

	fmt.Println("\n------------------------------------")
	fmt.Printf("Matrícula: %d\n", matricula)
	fmt.Printf("Salário Bruto:   R$ %.2f\n", salarioBruto)
	fmt.Printf("(-) Desconto INSS: R$ %.2f\n", inss)
	fmt.Printf("(-) Desconto IR:   R$ %.2f\n", ir)
	fmt.Println("------------------------------------")
	fmt.Printf("SALÁRIO LÍQUIDO: R$ %.2f\n", salarioLiquido)
	fmt.Println("------------------------------------")
}