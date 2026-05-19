package main
import "fmt"
func main() {
var C int
	fmt.Scan(&C)

	passageiros := 0

	for {
		var M int
		fmt.Scan(&M)

		passageiros += M

		if passageiros >= 2 * C {
			fmt.Println("hora de partir")
			break
		} else if passageiros == 0 {
			fmt.Println("vazio")
		} else if passageiros < C {
			fmt.Println("ainda cabe")
		} else {
			fmt.Println("lotado")
		}
	}
}

