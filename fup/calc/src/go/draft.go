package main
import "fmt"
func main() {

    var n1,n2 int
    var car string 

    fmt.Scan(&n1,&n2)
    fmt.Scan(&car)

    switch car {
    case "+":
        fmt.Println(n1 + n2)
    case "-":
        fmt.Println(n1 - n2)
    case "/":
        if  n2 != 0{
            fmt.Println(n1 / n2)
        }
    default:
        fmt .Println(n1 * n2)
    }

}
