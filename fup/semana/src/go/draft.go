package main
import "fmt"
func main() {
    var dia, hora int

    fmt.Scan(&dia, &hora)

    if dia == 1 {
        fmt.Println("NAO")
    } else if dia >= 2 && dia < 7{
        if hora >= 8 && hora < 12 || hora >= 14 && hora < 18{
            fmt.Println("SIM")
        } else if hora > 11 && hora < 14 {
            fmt.Println("NAO")
        } else{
            fmt.Println("NAO")
        }
    } else if dia == 7{
        if hora >= 8 && hora < 12 {
            fmt.Println("SIM")
        } else {
            fmt.Println("NAO")
        }
    }
}
