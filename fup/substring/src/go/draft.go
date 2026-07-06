package main
import "fmt"
func main() {
    var txt string
    var init, quant int

    fmt.Scan(&txt, &init, &quant)

    if init < 0 || init >= len(txt) || quant <= 0{
        fmt.Println("")

        return
    }

    fim := init + quant

    if fim > len(txt) {
        fim = len(txt)
    }
        fmt.Println(txt[init:fim])
}