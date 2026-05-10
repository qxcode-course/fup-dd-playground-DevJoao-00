package main
import "fmt"
func main() {
    var n int

    fmt.Scan(&n)

    if n % 2 == 0{
        return
    }else {
        for i := 0; i <= n; i++{
            if i % 2 == 1{
                fmt.Println(i)
            }
        }
        for a := n; a >= 0; a--{
            if a % 2 == 0{
                fmt.Println(a)
            }
        }
    }
}
