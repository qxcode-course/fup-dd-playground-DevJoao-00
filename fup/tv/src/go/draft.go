package main
import "fmt"
func main() {
	var valor float64
	var parcelas int

	fmt.Scan(&valor)
	fmt.Scan(&parcelas)

	juros := float64((parcelas - 1) * 5)

	total := valor * (1 + juros/100)
	parcela := total / float64(parcelas)

	fmt.Printf("%.2f\n", parcela)
	fmt.Printf("%.2f\n", total)
}