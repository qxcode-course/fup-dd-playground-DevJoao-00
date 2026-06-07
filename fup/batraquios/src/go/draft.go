package main
import "fmt"
func main() {
	var n, m int

	fmt.Scan(&n)

	vet1 := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&vet1[i])
	}

	fmt.Scan(&m)

	vet2 := make([]int, m)

	for i := 0; i < m; i++ {
		fmt.Scan(&vet2[i])
	}

	for i := 0; i < n; i++ {
		encontrou := false

		for j := 0; j < m; j++ {
			if vet1[i] == vet2[j] {
				encontrou = true
				break
			}
		}

		if encontrou == false{
			fmt.Println("nao")
			return
		}
	}

	fmt.Println("sim")
}