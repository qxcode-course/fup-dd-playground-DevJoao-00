package main
import (
    "fmt"
    "os"
    "bufio"
    "strings"
)
func main() {
    reader := bufio.NewReader(os.Stdin)

    texto, _ := reader.ReadString('\n')
    antiga, _ := reader.ReadString('\n')
    nova, _ := reader.ReadString('\n')

    texto = strings.TrimSpace(texto)
	antiga = strings.TrimSpace(antiga)
	nova = strings.TrimSpace(nova)

    for i := 0; i < len(texto); {
        if i <= len(texto) - len(antiga) && texto[i:i + len(antiga)] == antiga{
            fmt.Print(nova)
            i += len(antiga)
        } else {
            fmt.Printf("%c", texto[i])
            i++
        }
    }
        fmt.Println("")
}