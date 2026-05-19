package main
import "fmt"

func menorValor(vet[] int) int {
    menor := vet[0]

    for i:= 1; i < len(vet); i++{
        if vet[i] < menor{
            menor = vet[i]
        }
    }
        return menor
}
func main() {
    var a,b,c,d,e int

    fmt.Scan(&a, &b, &c, &d, &e)

    vet := []int {a,b,c,d,e}

    for i:= 0; i < 5; i++{
        fmt.Scan(&vet[i])
    }

    fmt.Println(menorValor(vet))
}
