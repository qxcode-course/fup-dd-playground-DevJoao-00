package main
import "fmt"
func main() {

    var p1, p2, p3 int
    var v1, v2, v3 float32 
    var qtddiero float32

    fmt.Scan(&p1, &p2, &p3)
    fmt.Scan(&v1, &v2, &v3)
    fmt.Scan(&qtddiero)

    total:= float32(p1) * v1 + float32(p2) * v2 + float32(p3) * v3 
    troco := qtddiero - total
    
    fmt.Printf("%.2f\n", troco)
}
