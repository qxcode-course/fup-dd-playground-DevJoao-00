package main
import "fmt"
func main() {
    var N int

    fmt.Scan(&N)

    pares := make([]int, N)

    cont := 0
    for i := 0; i < N; i++ {
        fmt.Scan(&pares[i])

        pares = append(pares, pares[i])

        if pares[i] % 2 == -1 {
            cont ++
        }

    }
        fmt.Println(cont)
}