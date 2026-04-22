package main
import "fmt"
func main() {
    var a,b,c int
    var h,l int

    fmt.Scan(&a,&b,&c,&h,&l)

    areaC := c * b 
    areaJ := h * l

    if areaC <= areaJ {
        fmt.Println("S")
    } else if areaC > areaJ {
        fmt.Println("N")
    }
}
