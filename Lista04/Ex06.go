package main
import "fmt"

func main() {
    var arr []int
    
  for i := 100; i > 0; i--{
      arr = append(arr, i)
  }
  
  for i := range arr{
      fmt.Println(arr[i])
  }
}
