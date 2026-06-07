package main
import "fmt"
func main() {
    var h,m,d,m2,a int

    fmt.Scan(&h,&m,&d,&m2,&a)

    if h < 10{
        fmt.Print("0")
    }
        fmt.Print(h)
        fmt.Print(":")

    if m < 10 {   
        fmt.Print("0")
    }
        fmt.Printf("%d ",m)

    if d < 10{
        fmt.Print("0")
    }   
        fmt.Print(d,"/")

    if m2 < 10{
        fmt.Print("0")
    }   
        fmt.Print(m2,"/")

        y := a % 100
        if y < 10{
            fmt.Print("0")
        }
        fmt.Println(y)

}