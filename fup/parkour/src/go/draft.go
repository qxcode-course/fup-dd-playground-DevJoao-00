package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    anterior := 0
    var x int
    parkour := 0
    for i := 1; i < n; i++ {
        fmt.Scan(&x)
        if x-anterior > 1 || x-anterior < -1 {
            anterior = x
            parkour++
        } else {
            anterior = x
            continue
        }
    }

    fmt.Println(parkour)

    var elementos int

    fmt.Scan(&elementos)

    vetorElementos := make([]int, elementos)

    for i := 0; i < elementos; i++ {
        fmt.Scan(&vetorElementos[i])
    }

    movimentos := 0

    for i := 0; i < elementos-1; i++ {
        if vetorElementos[i] < vetorElementos[i+1] && vetorElementos[i+1]-vetorElementos[i] > 1 {
            movimentos++
        } else if vetorElementos[i] > vetorElementos[i+1] && vetorElementos[i]-vetorElementos[i+1] > 1 {
            movimentos++
        }
    }

}