package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)

    var list[]int = make([]int, N)

    for i := range list{      
        fmt.Scan(&list[i])
    }

    for _, valor := range list{
        fmt.Println(valor)
    }
        if N == 0{
            fmt.Println("")
        }
}
