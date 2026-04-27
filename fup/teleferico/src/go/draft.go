package main
import "fmt"
func main() {
    var c,a int

    fmt.Scan(&c, &a)

    if c == a + 1 || c > a + 1{
        fmt.Println("1")
    } else if c / a + 1 < a{
        fmt.Println("2")
    } 
    
    
}
