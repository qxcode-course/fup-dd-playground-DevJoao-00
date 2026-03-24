package main
import "fmt"
func main() {
    var p1 int
    var p2 int
    var p3 int 
    var v1 float32
    var v2 float32 
    var v3 float32 
    var qtddiero float32
    fmt.Scan(&p1)
    fmt.Scan(&p2)
    fmt.Scan(&p3)
    fmt.Scan(&v1)
    fmt.Scan(&v2)
    fmt.Scan(&v3)
    fmt.Scan(&qtddiero)

    var sobra float32 = qtddiero - float32(p1) * v1 + float32(p2) * v2 + float32(p3) * v3 
    fmt.Printf("%.2f\n", sobra)
}
