package main
import "fmt"
func main() {
    
    var a,b int

    fmt.Scan(&a,&b)
        fmt.Printf("[ ")
    for i := a; i <= b -1; i++{
        fmt.Printf("%d ",i)
    }
        fmt.Println("]")
}