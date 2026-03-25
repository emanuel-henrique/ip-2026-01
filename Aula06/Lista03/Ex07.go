package main
import "fmt"

func main() {
  var a,b,c int
  fmt.Print("Digite o primeiro numero: ")
  fmt.Scan(&a)
  
  fmt.Print("Digite o segundo numero: ")
  fmt.Scan(&b)
  
  fmt.Print("Digite o terceiro numero: ")
  fmt.Scan(&c)
  
  var nums []int
  nums = append(nums, a,b,c)
  var numsOrganized [3]int
  
  for i,n := range nums {
    if i == 0 {
        if n > nums[i+1]{
            if n > nums[i+2]{
                numsOrganized[2] = n
            }else {
                numsOrganized[1] = n
            }
        }
        
        if n > nums[i+2]{
            numsOrganized[1] = n
        }
        numsOrganized[0] = n
    }
    
    if i == 1 {
        if n > nums[i-1]{
            if n > nums[i+1]{
                numsOrganized[2] = n
            }else{
                numsOrganized[1] = n
            }
        } else if n > nums[i+1]{
            numsOrganized[1] = n
        } else if  n < nums[i-1] &&  n < nums[i+1]{
            numsOrganized[0] = n
        }
    }
    
    if i == 2 {
        if n > nums[i-2]{
            if n > nums[i-1]{
                numsOrganized[2] = n
            }else{
                numsOrganized[1] = n
            }
        } else if n > nums[i-1]{
            numsOrganized[1] = n
        } else if  n < nums[i-2] &&  n < nums[i-1]{
            numsOrganized[0] = n
        }
    }
  }
  
  for i,_ := range numsOrganized{
    fmt.Println(numsOrganized[i])
  }
 
}
