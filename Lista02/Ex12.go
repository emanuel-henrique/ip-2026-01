package main
import "fmt"

func main() {
    var idade int
    fmt.Println("Digite sua idade: ")
    fmt.Scan(&idade)
    
    if idade > 0 && idade <= 2 {
        fmt.Println("Você é Recém-nascido")
    } else if idade >= 3 && idade <= 11 {
        fmt.Println("Você é Criança")
    } else if idade >= 12 && idade <= 19 {
        fmt.Println("Você é Adolescente")
    } else if idade >= 20 && idade <= 55 {
        fmt.Println("Você é Adulto")
    } else if idade > 55{
        fmt.Println("Você é Idoso")
    }
}
