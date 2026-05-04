package main
import "fmt"

func main() {
  length := 10
  var nums [10]int
  
  fmt.Println("Digite 10 numeros inteiros: ")
  
  var maiorNumero int
  var indiceMaior int
  
  for i := 0; i < length; i++{
      var num int
      fmt.Scan(&num)
      nums[i] = num
      if num > maiorNumero{
          maiorNumero = num
          indiceMaior = i
      }
  }
  
  fmt.Print(maiorNumero, indiceMaior)
}
