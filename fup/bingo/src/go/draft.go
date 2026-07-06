package main
import "fmt"

func main() {

	cartela := [4][4]int{
		{1, 9, 27, 23},
		{34, 20, 37, 47},
		{30, 87, 55, 69},
		{13, 60, 99, 66},
	}

	var numeros [6]int

	for i := 0; i < 6; i++ {
		fmt.Scan(&numeros[i])
	}

	acertos := 0

	for i := 0; i < 6; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 4; k++ {
				if numeros[i] == cartela[j][k] {
					acertos++
				}
			}
		}
	}

	fmt.Println(acertos)
}