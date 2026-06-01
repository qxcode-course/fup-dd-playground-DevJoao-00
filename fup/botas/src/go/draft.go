package main
import "fmt"
func main() {
    var n int

    fmt.Scan(&n)

    M := make([]int, n)
    L := make([]string, n)
    pares := 0
    for i:=0; i < len(M); i++{
        fmt.Scan(&M[i])
        fmt.Scan(&L[i])
            if len(M) == 3 {
                pares ++
                
    }
        fmt.Println(pares)
}
}
