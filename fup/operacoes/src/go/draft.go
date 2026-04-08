package main
import "fmt"
func main() {
    var n1 int 
    var n2 int

    fmt.Scan(&n1, &n2)

    soma := n1 + n2
    sub := n1 - n2
    mult := n1 * n2
    div := n1 / n2
    rest := n1 % n2

    fmt.Println(soma)
    fmt.Println(sub)
    fmt.Println(mult)
    fmt.Printf("%.2f\n", float64(div))
    fmt.Println(rest)
}
