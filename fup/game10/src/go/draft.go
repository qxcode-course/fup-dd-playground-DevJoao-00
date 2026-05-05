package main
import "fmt"
func main() {
    var N, D, A int

    fmt.Scan(&N, &D, &A)

    if N < 3 || N > 100{
        return
    }

    if D > A {
        A = (A + 1) % N
    }
    fmt.Println(A)
}
