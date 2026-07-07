package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
)
func main() {
    reader := bufio.NewReader(os.Stdin)

    txt, _:= reader.ReadString('\n')
    txt = strings.TrimSpace(txt)

    palavras := strings.Fields(txt)

    for i := 0; i < len(palavras) - 1; i++{
        if palavras[i] > palavras[i+1]{

            fmt.Println("nao")
        return
        }
    }
        fmt.Println("sim")

}