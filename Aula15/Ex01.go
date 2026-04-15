package main
import "fmt"

func main() {
  fmt.Println(potenciacao(6,1))
}

func potenciacao(x,n int) int {
    if n == 0{
        return 1
    } 
    return x * potenciacao(x,n-1)
}
