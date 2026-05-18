package main
import "fmt"
    
func maisvendido(contagem []int) string {
	if contagem[0] > contagem[1] {
		return "c"
	} else if contagem[1] > contagem[0] {
		return "l"
	}

	return "empate"
}

func turnoMaisVago(contagem []int) string {
	if contagem[0] < contagem[1] {
		return "m"
	} else if contagem[1] < contagem[0] {
		return "t"
	}

	return "empate"
}

func main() {
	var d int
	fmt.Scan(&d)

	sabores := make([]int, 2)
	turnos := make([]int, 2)

	for i := 0; i < d; i++ {
        
		var sabor, turno string
		fmt.Scan(&sabor, &turno)

		if sabor == "c" {
			sabores[0]++
		} else {
			sabores[1]++
		}

		if turno == "m" {
			turnos[0]++
		} else {
			turnos[1]++
		}
	}

	fmt.Println(maisvendido(sabores))
	fmt.Println(turnoMaisVago(turnos))
}

