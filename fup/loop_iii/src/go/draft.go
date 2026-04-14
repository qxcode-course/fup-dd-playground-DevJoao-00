package main
import "fmt"
func main() {
    var A, B int
    fmt.Scan(&A, &B)
        fmt.Printf("[ ")
    if A >= B{
        for i := A; i >= B + 1; i--{
            fmt.Printf("%d ",i)
        }
    }
        fmt.Println("]")
}
