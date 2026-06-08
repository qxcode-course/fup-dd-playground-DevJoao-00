package main
import "fmt"

func Fib(seq[]int, n int) int {

    vet := make([]int, n)

    vet[0] = 1
    vet[1] = 1

    for i := 2; i < n; i++{

        vet[i] = vet[i - 1] + vet[i + 1]

    }   
        return vet[n]
}
func main() {
    var n int

    fmt.Scan(&n)

    fmt.Println(Fib)
}