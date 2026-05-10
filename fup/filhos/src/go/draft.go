package main
import "fmt"
func main() {
    var fn, qf int

    fmt.Scan(&fn, &qf)
        if fn % 2 == 0 {
            for a := 1; a <= qf; a++{
            fmt.Println(fn)
             fn = fn + 2
        }
    }else {
        for a := 1; a <= qf; a++{
            fmt.Println(fn)
             fn = fn + 2
    }
    }
}
