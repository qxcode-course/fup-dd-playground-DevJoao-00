package main
import "fmt"
func main() {
    var player int
    var nd1,nd2 int

    fmt.Scan(&player)
    fmt.Scan(&nd1, &nd2)

    if player == 0 && (nd1 + nd2) % 2 == 0 {
        fmt.Println("0")
    } else if player == 1 && (nd1 + nd2) % 2 == 0 {
        fmt.Println("1")
    } else if player == 0 && (nd1 + nd2) % 2 != 0 {
        fmt.Println("1")
    } else {
        fmt.Println("0")
    }
}
