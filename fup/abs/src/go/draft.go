package main
import "fmt"
func main() {
    var n1 int
    var n2 int

    fmt.Scan(&n1, &n2)
    if n1 - n2 > 0{
        fmt.Println(n1 - n2)
    } else{
        fmt.Println(n2 - n1)
    }
}