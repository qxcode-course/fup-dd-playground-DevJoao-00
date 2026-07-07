package main
import "fmt"
func main() {
    var figmax, figatual int

    fmt.Scan(&figmax, &figatual) 

    repet := make([]int, 0)
    album := make([]bool, figmax + 1)

      var ant int

      for i := 0; i < figatual; i++{
        var x int
        fmt.Scan(&x)

        if i > 0 && x == ant{
            repet = append(repet, x)
        } else{
            album[x] = true
        }
            ant = x
      }
        fmt.Print("[ ")

    if len(repet) > 0{
        fmt.Print("")
        for _, v := range repet{
            fmt.Printf("%d ",v)
        }
    }
            fmt.Println("]")


            fmt.Print("[ ")

            vazio := false

            for i := 1; i <= figmax; i++{
                if !album[i]{
                    if !vazio{
                    fmt.Print("")
                    vazio = true
                }
                    fmt.Printf("%d ", i)
            }
        }
                fmt.Println("]")
}