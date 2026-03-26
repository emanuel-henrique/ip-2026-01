package main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite o numero: ")
	fmt.Scan(&num)
	
	if(num > 20 && num < 90){
	    fmt.Print("Este numero está entre 20 e 90")
	} else {
	    fmt.Print("Este numero não está entre 20 e 90")
	}
}
