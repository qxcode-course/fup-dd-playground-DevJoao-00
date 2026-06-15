package main
import "fmt"
func main() {
    var troco float64

    fmt.Scan(&troco)

    troco = float64(troco * 100) / 100.0

    vet := []float64{100.00, 50.00, 20.00, 10.00, 5.00, 2.00, 1.00, 0.50, 0.25, 0.10, 0.05}
        
    for _, moeda := range vet{
        qtd := int(troco / moeda)

        if qtd > 0 {
            fmt.Printf("%d de %.2f\n", qtd, moeda)
            troco = troco - float64(qtd) * moeda
        }
    }

    if troco > 0 {
        fmt.Printf("Falta %.2f\n", troco)
    }
}