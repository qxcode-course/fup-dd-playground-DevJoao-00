package main
import "fmt"
func main() {
    var A,B int

    fmt.Scan(&A, &B)

    menor := A

    fmt.Print("[ ")

    for B >= menor{
        fmt.Printf("%d %d ", A, B) 
        A++
        B--      
    }
        fmt.Println("]")
}
                

      

