package main
import "fmt"
func main() {
    var n int

    fmt.Scan(&n)

    v := make([]int, n)

    maior := 0

    for i := 0; i < n; i++{
        fmt.Scan(&v[i])
            if v[i] > maior{
                maior = v[i]
            }
    }

    for line := maior; line >= 1; line --{
        for i := 0; i < n; i++{
            if v[i] >= line{
                fmt.Print("#")
            }else{
                fmt.Print("_")
            }
        }
            fmt.Println("")
    }
}