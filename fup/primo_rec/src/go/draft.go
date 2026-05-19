package main
import "fmt"
func ehPrimo(n int, div int) int {
    // números menores que 2 não são primos
    if n < 2 {
        return 0
    }

    // terminou os testes
    if div == 1 {
        return 1
    }

    // encontrou divisor
    if n % div == 0 {
        return 0
    }

    // chamada recursiva
    return ehPrimo(n, div-1)
}

func main() {
    var N int

    fmt.Scan(&N)

    fmt.Println(ehPrimo(N, N-1))
}
