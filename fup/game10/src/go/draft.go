package main
import "fmt"
func main() {
    var N, D, A int

    fmt.Scan(&N, &D, &A)

    if N < 3 || N > 100{
        return
    }
    if D == A{
        A = (D - A)
    } else if D > A {
        A = (D - A) % N
    } else if D < A {
        A = (D - A % N + N) % N
    }
        fmt.Println(A)
}
