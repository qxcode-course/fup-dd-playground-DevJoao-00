package main
import (
    "fmt"
)
type info struct{
    v int
    c int
    p int
}
func contv(texto string) int{
    cont := 0

    for i := 0; i < len(texto); i++{
        if texto[i] == 'a' || texto[i] == 'e' || texto[i] == 'i' ||texto[i] == 'o' || texto[i] == 'u'{
            cont ++
        }
    }
    return cont
}
func analisar(texto string) info{
    var pal info

    return pal
}
func main() {
    fmt.Println("Hello, World!")
}