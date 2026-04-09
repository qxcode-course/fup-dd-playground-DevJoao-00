package main
import "fmt"
func main() {
    var v1, v2, v3 int

    fmt.Scan(&v1, &v2, &v3)

    if v1 > v2 && v1 < v3{
        fmt.Println(v1)
    } else if v2 > v1 && v2 < v3 {
        fmt.Println(v2)
    } else {
        fmt.Println(v3)
    }
}
