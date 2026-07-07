package main
import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func vogal (x byte) bool{
    return x == 'a' || x == 'e' || x == 'i' || x == 'o' || x == 'u'
} 

func main() {
    reader := bufio.NewReader(os.Stdin)

    frase, _ := reader.ReadString('\n')
    frase = strings.TrimSpace(frase)

    sep := strings.Fields(frase)

    fmt.Print(sep[0])

    for i := 1; i < len(sep); i++{
        if vogal(sep[i-1][len(sep[i-1])-1]) && vogal(sep[i][0]) {
			j := 0
			for j < len(sep[i]) && vogal(sep[i][j]) {
				j++
			}
			if j > 0 {
				fmt.Print(sep[i][j-1:])
			} else {
				fmt.Print(sep[i])
			}
		} else {
			fmt.Print(" ", sep[i])
		}
	}
	    fmt.Println("")
    
}