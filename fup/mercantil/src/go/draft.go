package main
import "fmt"
func main() {
    var rodada int

    fmt.Scan(&rodada)

    produtos := make([]float64, rodada)
    chute := make([]float64, rodada)
    escolhas := make([]string, rodada)

        for i := 0; i < rodada; i++{
            fmt.Scan(&produtos[i])
        }

        for i := 0; i < rodada; i++{
            fmt.Scan(&chute[i])
        }
        
        for i := 0; i < rodada; i++{
            fmt.Scan(&escolhas[i])
        }

        jog1 := 0
        jog2 := 0

            for i := 0; i < rodada; i++{
                if chute[i] == produtos[i]{
                    jog1++
                        continue
                }

                if escolhas[i] == "M"{
                    if produtos[i] > chute[i]{
                        jog2++
                    } else {
                        jog1++
                    }
                }
                if escolhas[i] == "m"{
                    if produtos[i] < chute[i]{
                        jog2++
                    }else {
                        jog1++
                    }
                }
            }

            if jog1 > jog2{
                fmt.Println("primeiro")
            } else if jog1 < jog2 {
                fmt.Println("segundo")
            } else{
                fmt.Println("empate")
            }
}
