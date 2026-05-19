package main
import "fmt"
func main() {
    var N int
   
    fmt.Scan(&N)

    inteiros := make([]int, N) 

    for i := 0; i < N; i ++{
        fmt.Scan(&inteiros[i])
    }

    fmt.Print("[ ")

    for i := 0; i < N; i ++{
        fmt.Print(inteiros[i], " ")
    }

    fmt.Println("]")
    
}
