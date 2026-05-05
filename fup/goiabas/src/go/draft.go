package main
import "fmt"
func main() {
    var C, b, g, m  int

    fmt.Scan(&C, &b, &g, &m)

    frutas := b + g + m

    if frutas <= C {
        fmt.Println("1")
    } else {
        tempo := frutas / C

        if frutas % C != 0{
            tempo++
        }
            fmt.Println(tempo)
    }
}