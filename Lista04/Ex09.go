package main
import (
    "fmt"
)

func main() {
    var alturas []float64
    var acimaDaMedia []float64
    var sum float64
    
    for range 10{
        var altura float64
        fmt.Scan(&altura)
        alturas = append(alturas, altura)
    }
    
    for _,n := range alturas{
        sum += n
    }
    
    media := sum/10
    
    for _,n := range alturas{
        if n > media{
            acimaDaMedia = append(acimaDaMedia, n)
        }
    }
    
    fmt.Print(acimaDaMedia)
   
}
