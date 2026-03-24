package main
import "fmt"
func main() {

    var a, b int 

    fmt.Scan(&a, &b) 

    var media = float32(a + b) / 2

    fmt.Printf("%.1f\n", media)
}
