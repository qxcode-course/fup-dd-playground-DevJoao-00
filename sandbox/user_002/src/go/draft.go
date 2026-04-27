package main
import (
    "fmt"
    "slices"
)
func filtrar_impares(nums[]int) []int{
    impares := make([]int, 0, len(nums)) 
    for _, elem := range nums{
        if elem % 2 == 1{
            impares = append(impares, elem)
        }
    }
            return impares
}
func index(nums[] int, valor int) int{
    for i, elem := range nums {
        if elem == valor{
            return i
        }
    }
            return -1
}
func contain(nums[]int, valor int) bool{
    for _, elem := range nums{
        if elem == valor{
            return true
        }
    }
            return false
}
func count(nums[]int, valor int) int{
    contador := 0
    for _, elem := range nums{
        if elem == valor{
            return contador + 1
        }
    }
            return contador
}
func separar_figs(montante []int)([]int, []int){
    album := make([]int, 0, len(montante))
    repet := make([]int, 0, len(montante))
        for _, figs := range montante{
            if contain (album, figs){
            album = append(album,figs)
        } else {
            album = append(repet, figs)
        }
    }
            return album, repet
}
func main() {
    var montante []int = make([]int, 0, 1)
    fmt.Println(montante, len(montante), cap(montante))

	album, repet := separar_figs(montante)

	slices.Sort(repet)
	fmt.Println(album)
	fmt.Println(repet)
    
}
