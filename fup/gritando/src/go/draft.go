package main
import (
    "fmt"
    "bufio"
    "os"
)
func main() {
     reader := bufio.NewReader(os.Stdin)

    frase, _ := reader.ReadString('\n')

    for _, i := range frase{
        if i >= 'a' && i <= 'z'{
            fmt.Printf("%c",i - 32)
        } else if i >= 'A' && i <= 'Z'{
            fmt.Printf("%c", i + 32)
        }else {
            fmt.Printf("%c",i)
        }
    }
}