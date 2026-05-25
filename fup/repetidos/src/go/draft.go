package main
import "fmt"
func main() {
    var p, n int

    fmt.Scan(&p, &n)

    vet := make([]int, n)
    repet := 0
    for i := 0; i < n; i++ {
        fmt.Scan(&vet[i])

        if vet[i] == p{
            repet ++
        }
    }
        fmt.Println(repet)
}
