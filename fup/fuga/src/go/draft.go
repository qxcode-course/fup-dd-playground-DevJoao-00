package main
import "fmt"
func main() {
    var H, P, F, D int

    fmt.Scan(&H, &P, &F, &D)

    dif := H - P

    if H >= 0 && H < 16 && P >= 0 && P < 16 && F >= 0 && F < 16 && H != P && H != F && P != H{
       for H != F{
            if F > dif {
                F = (F - D) % 16
                fmt.Println("S") 
                break
            } 

        }
            if F < dif {
                F = P
                fmt.Println("N")
                
        }   
    } else {
        return
    }
}
