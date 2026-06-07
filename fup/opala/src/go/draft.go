package main

import "fmt"

func main() {
    var v, t int
    var c float64

    fmt.Scan(&v, &t, &c)

    dist := float64(v) * (float64(t) / 60.0)
    desemp := dist / c

    fmt.Printf("%.2f\n", desemp)
}