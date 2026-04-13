package main

import (
	"fmt"
	"math"
)
func main() {
    var a,b,c float64

    fmt.Scan(&a, &b, &c)

    delta := (math.Pow(b, 2) - 4 * a * c)

    if delta < 0 {
        fmt.Println("nao ha raiz real")
    } else if delta == 0 {
        x := (-b + math.Sqrt(delta)) / (2 * a) 
        fmt.Printf("%.2f\n", x)
    } else if delta > 0 {
        x1 := (-b + math.Sqrt(delta)) / (2 * a)
        x2 := (-b - math.Sqrt(delta)) / (2 * a)
        fmt.Printf("%.2f\n", x1)
        fmt.Printf("%.2f\n", x2)
    }
}
