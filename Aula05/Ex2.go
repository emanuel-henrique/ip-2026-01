package main
import "fmt"

func main() {
  var nt1, p1 float64
  var nt2, p2 float64
  
  fmt.Print("Digite a primeira nota e seu peso nesta ordem: ")
  fmt.Scan(&nt1, &p1)
  
  fmt.Print("Digite a segunda nota e seu peso nesta ordem: ")
  fmt.Scan(&nt2, &p2)
  
  mediaPound := ((nt1 * p1) + (nt2 * p2)) / (p1 + p2)
  fmt.Printf("A média ponderada é: %.2f", mediaPound)
}
