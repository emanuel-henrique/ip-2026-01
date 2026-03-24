package main
import "fmt"

func main() {
    var l1,l2,l3 int
    fmt.Print("Digite o valor de cada lado do triangulo (Ex: 10 20 10): ")
    fmt.Scan(&l1,&l2,&l3)
    
    if l1 + l2 > l3 && l1 + l3 > l2 && l3 + l2 > l1{
        if l1 == l2 && l2 == l3 {
            fmt.Print("Este é um triângulo equilatero!")
        } else if l1 != l2 && l1 != l3 && l2 != l3 {
            fmt.Print("Este é um triângulo escaleno!")
        } else {
            fmt.Print("Este é um triângulo isóceles!")
        }
    } else {
        fmt.Print("Este triangulo não é válido")
    }
}
