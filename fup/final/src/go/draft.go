package main
import "fmt"
func main() {
    var nota1, nota2, notaf int

    fmt.Scan(&nota1, &nota2, &notaf)

    if (nota1 + nota2) / 2  >= 7 {
        fmt.Println("aprovado")
    } else if  (nota1 + nota2) / 2 < 7 {
        media := int(nota1 + nota2) + notaf / 2

        if media >= 5 {
        fmt.Println("aprovado na final")
        }   else {
        fmt.Println("reprovado na final")
        }        

    }
}