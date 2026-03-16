package main

import "fmt"

// Conversão de temperatura
func main() {
    var count int

    fmt.Print("Quantas temperaturas deseja converter? ")
    fmt.Scan(&count)

    var temperatures []float64
    var convertedTemperatures []float64

    for i := 1; i <= count; i++{
        var temperature float64
        fmt.Printf("Digite a %dº temperatura (FAHRENHEIT): ", i)
        fmt.Scan(&temperature)
        temperatures = append(temperatures, temperature)
    }

    fmt.Println("===================================================")

    for _, n := range temperatures{
        convertedTemperature := convertTemperature(n)
        convertedTemperatures = append(convertedTemperatures, convertedTemperature)
        fmt.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS\n", n, convertedTemperature)
    }
}

func convertTemperature(temperature float64) float64{
    convertedTemperature := ((temperature - 32) * 5) / 9
    return convertedTemperature
}