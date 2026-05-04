package main
import "fmt"

func main() {
    var arr []int
    var impares []int
    
    
  for i := 100; i > 0; i--{
      arr = append(arr, i)
      if i%2 != 0{
        impares = append(impares, i)
      }
  }
  
  for i := range impares{
      fmt.Println(impares[i])
  }
}
