package main
import "fmt"

func desp(Pessoas[]int)int{
    cont := 0

    for i := 0; i < len(Pessoas); i++{
        if i == 0 && Pessoas[i] == 0 && Pessoas[i - 1] == 0{
            cont ++
        }
    }   
        return cont
}
func main() {
    var n int

    fmt.Scan(&n)
    vet := make([]int, n)

    for i:= 0; i < n; i++{
    fmt.Scan(&vet[i])


    }
    fmt.Println(desp(vet))
}