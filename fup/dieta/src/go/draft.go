package main
import "fmt"
func main() {
    var n int

    fmt.Scan(&n)

    cal := make([]int, n)
    soma := 0

    for i := 0; i < len(cal); i ++{
        fmt.Scan(&cal[i])
        soma += (cal[i]) 
    }

    media := float64 (soma) / float64(n)
        fmt.Printf("%.1f\n", media )
}

    
