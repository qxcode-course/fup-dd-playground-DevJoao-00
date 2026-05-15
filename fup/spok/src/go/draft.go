package main
import "fmt"

func main() {
    var ID int

    fmt.Scan(&ID)

    num := ID
    inverso := 0

    for{
        ud := ID % 10
        inverso = inverso * 10 + ud
        ID = ID / 10

        if ID == 0 {
            break
        }
        
    }
        if num == inverso{
            fmt.Println("1")
        }else {
            fmt.Println("0")
        }
}
