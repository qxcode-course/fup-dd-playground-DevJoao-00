package main
import "fmt"
func main() {
    var segundos int

    fmt.Scan(&segundos)

    horas := segundos / 3600
    minutos := (segundos % 3600) / 60
    segRest := segundos % 60

    fmt.Printf("%d:%d:%d\n", horas, minutos, segRest)
}
