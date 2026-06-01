package main
import "fmt"
func main() {
    var n1 int 
    var n2 int

    fmt.Scan(&n1, &n2)

    soma := n1 + n2
    sub := n1 - n2
    mult := n1 * n2
    div := float64(n1) / float64(n2)
    rest := n1 % n2

    fmt.Println(soma)
    fmt.Println(sub)
    fmt.Println(mult)
    fmt.Printf("%.2f\n", div)
    fmt.Println(rest)

}
