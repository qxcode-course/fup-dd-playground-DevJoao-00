package main
import "fmt"
func main() {
    var n int

    fmt.Scan(&n)

    fila := make([]int, n)
    
        for i := 0; i < n; i++{
            fmt.Scan(&fila[i])
        }
        fmt.Print("[ ")
        for _, i := range fila{
            if i % 2 != 0 {
                fmt.Printf("%d ", i)
            }
        }   
            fmt.Println("]")
            
            fmt.Print("[ ")
        for _, p := range fila{
            if p % 2 == 0 {
                fmt.Printf("%d ", p)
            }
        }
            fmt.Println("]")
}