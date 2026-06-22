package main
import "fmt"
func main() {
    var n int

    fmt.Scan(&n)

    vet := make([]int, n)

    zero := true

    for i:= 0; i < n; i++ {
        fmt.Scan(&vet[i])

         if vet[i] != 0{
           zero = false
        }
    }

    

    if zero {
        fmt.Println(vet[0])
        return
    }

    for _, v := range vet{
            fmt.Print(v)
    }
        fmt.Println("")
}