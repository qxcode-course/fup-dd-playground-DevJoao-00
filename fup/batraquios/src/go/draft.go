package main
import "fmt"
func main() {
    var n,m int

    fmt.Scan(&n, &m)

    vetn := make([] int, n)
    vetm := make([] int, m)

    for i := 0; i < n; i++ {
        fmt.Scan(&vetn[i])
    }
    
    for j := 0; j < m; j++ {
            fmt.Scan(&vetm[j])
        }
        fmt.Println(vetn)
        fmt.Println(vetm)
}
