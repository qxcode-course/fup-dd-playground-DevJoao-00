package main
import "fmt"
func main() {
    var ad,bd,cd int
    var ac,bc int

    fmt.Scan(&ad,&bd,&cd,&ac,&bc)

    if bd <= ac && cd <= bc{
        fmt.Println("S")
    } else if bd > ac && cd > bc{
        fmt.Println("N")
    } 
}
