package main
import "fmt"
import "strconv"

func main() {
    var num int 
    
    fmt.Print("Digite um número com 3 casa (Ex.: 103): ")
    fmt.Scan(&num)
    
    if num < 100 || num > 999 {
        fmt.Print("Número Inválido!")
        return
    }
    strNum := strconv.Itoa(num)
    
    fmt.Printf("Dezena: %c0", strNum[1])

}
