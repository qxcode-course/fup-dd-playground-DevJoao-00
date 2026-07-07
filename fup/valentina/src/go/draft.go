package main
import "fmt"
func main() {
    var c1, c2 rune
    var op string

    fmt.Scanf("%c\n", &c1)
    fmt.Scanln(&op)
    fmt.Scanf("%c", &c2)

    x := int(c1 - 'a')
    y := int(c2 - 'a')

    var resp int

    if op == "+"{
        resp = (x + y) % 26
    } else {
        resp = (x - y + 26) % 26
    }
        fmt.Printf("%c\n", rune(resp) + 'a')
}