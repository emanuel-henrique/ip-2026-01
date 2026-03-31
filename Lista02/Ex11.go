package main
import "fmt"

func main() {
    var x float64
    fmt.Print("Digite o valor de X: ")
    fmt.Scan(&x)
    if(x == 2){
         fmt.Print("Número Invalido")
         return
    }
    fmt.Print(calculeFdeX(x))
    
    
}

func calculeFdeX(x float64) float64{
    return 8 / (2-x)
}
