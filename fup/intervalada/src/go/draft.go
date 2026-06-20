package main
import "fmt"
func main() {
    var n, li, ls int 

    fmt.Scan(&n,&li,&ls)

    vet := make([]int, n)
    cont := 0

    for i := 0; i < n; i++ {
        fmt.Scan(&vet[i])

        if vet[i] >= li && vet[i] <= ls {
            cont++
        }
        
    }
        fmt.Println(cont)
}