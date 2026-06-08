package main
import "fmt"

func soma(a,b int) int{
    soma := 0

    if a > b{
        c := a
        a = b
        b = c
    }
    for i := a; i <= b; i ++{
        soma += i
    }
        return soma
}
func main() {
    var a, b int

    fmt.Scan(&a, &b)

    fmt.Println(soma(a,b))
}