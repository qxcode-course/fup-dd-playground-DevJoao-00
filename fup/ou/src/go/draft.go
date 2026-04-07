package main
import "fmt"
func main() {
    var n1 int
    fmt.Scan(&n1)

    if n1 == 3 || n1 == 5 {
        fmt.Println("SIM")
    } else {
        fmt.Println("NAO")
    }
}
