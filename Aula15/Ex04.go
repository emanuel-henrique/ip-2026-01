package main

import "fmt"

func main() {
    numero := 25
    fmt.Printf("O número %d em binário é: ", numero)
    decimalParaBinario(numero)
    fmt.Println()
}

func decimalParaBinario(n int) {
    if n == 0 {
        return
    }

    decimalParaBinario(n / 2)

    fmt.Printf("%d", n%2)
}
