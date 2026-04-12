package main
import "fmt"

func ceil(nota float64) int {
    inteiro := int(nota)

    if nota > 0 && nota != float64(inteiro){
        return inteiro + 1
    }
        return inteiro
}
func floor(nota float64) int {
    inteiro := int(nota)

    if nota < 0 && nota != float64(inteiro){
        return inteiro - 1
    }
        return inteiro
}
func round(nota float64) int {
    inteiro := int(nota)
    partdecimal := nota - float64(inteiro)

    if inteiro < 0 && partdecimal != 0{
        partdecimal = float64 (inteiro) - nota
    }
    if partdecimal >= 0.5 {
        return ceil(nota)
    } else {
        return floor(nota)
    }
        
}
func main() {
    var car string
    var nota float64

    fmt.Scan(&car, &nota)

    if car == "c" {
        fmt.Println(ceil(nota))
    } else if car == "f" {
        fmt.Println(floor(nota))
    } else if car == "r" {
        fmt.Println(round(nota))
    }
}
