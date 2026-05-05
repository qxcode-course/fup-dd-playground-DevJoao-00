package main
import "fmt"
func main() {
    var N, X, Y int 
    var C string
    var S int

    fmt.Scan(&N, &X, &Y, &C, &S)

        switch C {
            case "U":
                Y = (Y - S % N + N) % N
            case "R":
                X = (X + S) % N
            case "D":
                Y = (Y + S) % N    
            case "L":
                X = (X - S % N + N) % N
            }
                fmt.Println(X,Y)
    }

