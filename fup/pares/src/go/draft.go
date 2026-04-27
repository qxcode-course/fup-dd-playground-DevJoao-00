package main
import "fmt"
func main() {
     var a, b int

    fmt.Scan(&a,&b)

    if a > b {
        fmt.Print("invalido")
    } 
    soma := 0

    for i := a; i <= b; i++{
        if i % 2 == 0 {
           soma += i
        }
     }   
     if soma == 0 {
            fmt.Println("")
     } else {
         fmt.Println(soma)
     }           
}
