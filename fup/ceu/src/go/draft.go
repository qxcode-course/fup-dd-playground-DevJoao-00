package main
import "fmt"
func main() {
    
    var N int
    var n10 = "ceu" 

    fmt.Scan(&N)
        fmt.Print("[ ")
    for i:= 0; i <= 9; i++ {
        if i == N{
            continue
        }
        fmt.Printf("%d ",i)
    }
       if N != 10 {
         fmt.Println(n10,"]")        
       }else{
        fmt.Println("]")
       
}
}
