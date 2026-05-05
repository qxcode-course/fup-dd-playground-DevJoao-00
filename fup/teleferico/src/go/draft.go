package main
import "fmt"
func main() {
    var c,a int

    fmt.Scan(&c, &a)

    v := (a + 1) / c 

    if (a + 1) % c != 0{
        fmt.Println(v + 1)
    } else {
        fmt.Println(v + v - 1)
    } 
    
    
    
}
