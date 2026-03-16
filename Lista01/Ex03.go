package main

import (
	"fmt"
	"math"
	"strconv"
)

// Composição Inteira (+)
func main() {
    var n1, n2, n3 int
    
    for i := 1; i <= 3; i++ {
        var number int
        fmt.Print("Digite o primeiro numero: ")
        fmt.Scan(&number) 
   
        algarismsCount := len( fmt.Sprintf("%d", number))
        
        if algarismsCount > 1 {
            fmt.Print("DIGITO INVALIDO")
            return 
        }
        
        if i == 1 {
            n1 = number
        } else if i == 2 {
            n2 = number
        } else {
            n3 = number
        }
    }
    
    nt1 := fmt.Sprintf("%d%d%d", n1, n2, n3) 
    n, _ := strconv.Atoi(nt1)
    nt2 := math.Pow(float64(n), 2)
    fmt.Printf("%s, %d\n", nt1, int(nt2))
}