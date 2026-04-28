package main
import "fmt"
func main() {
    var a,b int

    fmt.Scan(&a, &b)

    if b >= a {
        for i := a; i <= b; i++{
            if i % 2 == 0 && i % 3 == 0{
                soma := i + i
                fmt.Print(soma)
            } 
        }
    } else if a > b{
        fmt.Println("invalido")
    }


}
