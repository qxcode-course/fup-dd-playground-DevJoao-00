package main
import "fmt"
func main() {
    var n1, n2 int
    fmt.Scan(&n1, &n2)

    for i:= n1; i < n2; i++{
        fmt.Println(i)
        if i % 3 == 0 {
            fmt.Println("zig")
        } else if i % 5 == 0{
            fmt.Println("zag")
        } else if i % 3 == 0 && i % 5 == 0{
            fmt.Println("zigzag")
        }
    }

}
