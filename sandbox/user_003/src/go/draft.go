package main
import "fmt"
func main() {
    var nome string = "João"
    var letras[]rune = []rune(nome)

    fmt.Println(len(nome))
    fmt.Println(len(letras))
}
