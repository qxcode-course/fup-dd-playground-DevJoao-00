package main
import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var arena [100][100]string

	linhaLeao := -1
	colunaLeao := -1

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&arena[i][j])

			if arena[i][j] == "L" {
				linhaLeao = i
				colunaLeao = j
			}
		}
	}

	pontosG := 0
	pontosC := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {

			
			if i == linhaLeao || j == colunaLeao {
				continue
			}

			switch arena[i][j] {
                case "G":
				    pontosG += 2
			    case "C":
				    if i+j == n-1 {
				    	pontosC += 2
				     }else{
					    pontosC++
				    }
			}
		}
	}

	if pontosG > pontosC {
		fmt.Println("Gladiadores")
	} else if pontosC > pontosG {
		fmt.Println("Condenados a morte")
	} else {
		fmt.Println("Ninguem")
	}
}