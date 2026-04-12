package main
import "fmt"
func main() {
    var a, b int

    fmt.Scan(&a, &b)

    for a := a; a < b; a++ {
        if a <= b{
        fmt.Println(a)
        } else {
            fmt.Println("Insira um valor valido")
        }
    }
}
