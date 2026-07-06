package main
import (
    "fmt"
    "bufio"
    "os"
)
func main() {
    leitor := bufio.NewReader(os.Stdin)

	frase, _ := leitor.ReadString('\n')

	if len(frase) > 0 && frase[len(frase)-1] == '\n' {
		frase = frase[:len(frase)-1]
	}

    x := []rune(frase)

    for i := len(x) - 1; i >= 0; i --{
        fmt.Print(string(x[i]))
    }

    fmt.Println("")
}
