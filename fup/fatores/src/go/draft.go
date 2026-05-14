package main
import "fmt"

func main() {

    var N int
    var f = 2
    var c = 0

    fmt.Scan(&N)

    for N != 1{
        if N % f == 0{
            N = N / f
            c++
        } else {
            if c > 0{
                fmt.Printf("%d %d\n", f, c)
                
            }
                f++
                c = 0        
        }
    }
        if c > 0 {
            fmt.Printf("%d %d\n", f, c)
        }
}
           
    

