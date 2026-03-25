package main
import "fmt"

func main() {
  var qA float64
  fmt.Print("Quantos alunos verificar: ")
  fmt.Scan(&qA)
  
  var qtN float64
  fmt.Print("Quantas notas serão avaliadas: ")
  fmt.Scan(&qtN)
  
  var total float64
  fmt.Print("=================================\n")
  for i := 1.0; i <= qA; i++ {
      for j := 1.0; j <= qtN; j++ {
          var nota float64
          fmt.Printf("Digite a %2.fº do %2.fº aluno: ", j,i)
          fmt.Scan(&nota)
          total += nota
      }
      fmt.Print("=================================\n")
  }
  fmt.Print("=================================\n")
  var totalNotas = float64(qA*qtN)
  var result  = total / totalNotas
  fmt.Printf("A media da turma é: %.1f", result)
}
