package main
import "fmt"
func main() {

    var N int
    
    fmt.Scan(&N)
        fmt.Print("[ ")
    for i:= 0; i <= 9; i++ {
        if i == N{
            continue
        }
        fmt.Printf("%d ",i)
    }
        fmt.Println("ceu ]")
}
