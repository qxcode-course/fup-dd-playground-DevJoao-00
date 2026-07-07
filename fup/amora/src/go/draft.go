package main
import (
    "fmt"
    "os"
    "bufio"
)
func main() {
    reader := bufio.NewReader(os.Stdin)

    texto, _ := reader.ReadString('\n')
    trecho, _ := reader.ReadString('\n')

    texto = texto[:len(texto)-1]
    trecho = trecho[:len(trecho)-1]

    cont := 0

    for i := 0; i <= 
    len(texto) - len(trecho); i++{
        if texto[i:i+len(trecho)] == trecho{
            cont++
        }
    }
        fmt.Println(cont)
}