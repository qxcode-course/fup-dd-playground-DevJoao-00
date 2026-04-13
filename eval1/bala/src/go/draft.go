package main
import "fmt"
import "math"
func main() {
    var x1,x2,y1,y2 float64

    fmt.Scan(&x1,&y1,&x2,&y2)

    DAB := math.Sqrt(math.Pow(x2 - x1, 2) + math.Pow(y2 - y1, 2))

    fmt.Printf("%.2f\n", DAB)
}
