package main

import (
	"bufio"
	"fmt"
	"os"
)

func vogal(c rune) bool{
    return c == 'a' || c == 'e'|| c == 'i'|| c == 'o'|| c == 'u'|| 
    c == 'A'|| c == 'E'|| c == 'I'|| c == 'O'|| c == 'U'
}

func main() {
	reader := bufio.NewReader(os.Stdin)

    frase, _ := reader.ReadString('\n')

    for _, i := range frase{
        if i == '\n'{
            break
        }

        if i == ' '{
            fmt.Print(" ")
        } else if vogal(i){
            fmt.Print("v")
        } else {
            fmt.Print("c")
        }
    }
        fmt.Println("")
}
