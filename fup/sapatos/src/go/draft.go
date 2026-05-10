package main
import "fmt"
func main() {
    var a,b int
    var soma = 0 

    fmt.Scan(&a, &b)

    if b >= a {
        for i := a; i <= b; i++{
            if i % 2 == 0 && i % 3 == 0{
                soma = soma + i
            } 
        }
    } else if a > b && soma == 0{
        fmt.Print("invalido")
    }
        if soma ==  0{
            fmt.Println("")
        }else{
            fmt.Println(soma)
        }
}
