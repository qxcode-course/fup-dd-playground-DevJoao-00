package main
import "fmt"
func main() {
    var n1 int 
    var n2 int
    
    fmt.Scan(&n1, &n2)

    div := n1 / n2
    rest := n1 % n2

    fmt.Println(div)
    fmt.Println(rest)
    fmt.Printf("%.2f\n", float64 (n1) / float64(n2))
}
