package main
import "fmt"
func main() {
    var a1, a2, a3 int
    fmt.Scan(&a1,&a2,&a3)

    if a1 == a2 && a1 == a3 && a2 == a3{
        fmt.Println("3")
    } else if a1 == a2 && a1 != a3 {
        fmt.Println("2")
    } else if a1 == a3 && a1 != a2 {
        fmt.Println("2")
    } else if a2 == a3 && a2 != a1 {
        fmt.Println("2")
    } else {
        fmt.Println("0")
    }

}
