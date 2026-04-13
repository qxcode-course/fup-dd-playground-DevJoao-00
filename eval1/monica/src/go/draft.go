package main
import "fmt"
func main() {
    var M, a, b int
    
    fmt.Scan(&M, &a, &b)

    IMV := M - (a + b)
    
    if a > b && a > IMV{
        fmt.Println(a)
    } else if b > a && b > IMV {
        fmt.Println(b)
    } else{
        fmt.Println(IMV)
    }    
}
