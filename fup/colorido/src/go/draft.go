package main
import "fmt"
func main() {
    var N int
    var pe string
    var n10 ="ceu"

    fmt.Scan(&N, &pe)

    fmt.Print("[ ")

    for i := 0; i < 10; i++{
        if i == N{
            continue
        }

         fmt.Printf("%d%s ", i, pe)

         if  pe == "e"{
            pe = "d"
         } else{
            pe = "e"
         }
    }

    if N != 10{
        fmt.Println(n10, "]")
    } else {
        fmt.Println("]")
    }
}
