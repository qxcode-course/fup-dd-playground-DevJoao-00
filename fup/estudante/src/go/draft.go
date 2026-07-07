package main
import (
	"bufio"
	"fmt"
	"os"
)
type Aluno struct {
	nome        string
	n1, n2, n3 float64
	media       float64
}

func main() {
	entrada := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(entrada, &n)
	entrada.ReadByte() 

	estudante := make([]Aluno, n)

	for i := 0; i < n; i++{
		nome, _ := entrada.ReadString('\n')
		nome = nome[:len(nome)-1]

		estudante[i].nome = nome

		fmt.Fscan(entrada, &estudante[i].n1, &estudante[i].n2, &estudante[i].n3)
		entrada.ReadByte()

		estudante[i].media = (estudante[i].n1 + estudante[i].n2 + estudante[i].n3) / 3
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if estudante[j].media < estudante[j+1].media {
				estudante[j], estudante[j+1] = estudante[j+1], estudante[j]
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d: %s\n", i, estudante[i].nome)
		fmt.Printf("   Media: %.2f\n", estudante[i].media)
		fmt.Printf("   N1: %.2f, N2: %.2f, N3: %.2f\n",
			estudante[i].n1,
			estudante[i].n2,
			estudante[i].n3)
	}
}