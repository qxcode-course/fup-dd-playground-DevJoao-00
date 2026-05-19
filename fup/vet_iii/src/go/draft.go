package main
import "fmt"
func imp(inteiros[] int){

    fmt.Print("[")
    
    for i := 0; i < len(inteiros); i ++{
        fmt.Print(inteiros[i])

        if i != len(inteiros) - 1{
            fmt.Print(", ")
        }
    }
     fmt.Println("]")
}
func main() {
      var N int
   
    fmt.Scan(&N)

    inteiros := make([]int, N) 

    for i := 0; i < N; i ++{
        fmt.Scan(&inteiros[i])
    } 
    imp(inteiros)
}

