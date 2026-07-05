package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	jogarNovamente := 1

	// jogadas[i] contém os índices que a jogada i vence
	// 0-Pedra 1-Papel 2-Tesoura 3-Lagarto 4-Spock
	vence := [][]int{
		{3, 2}, // Pedra vence Lagarto e Tesoura
		{0, 4}, // Papel vence Pedra e Spock
		{1, 3}, // Tesoura vence Papel e Lagarto
		{4, 1}, // Lagarto vence Spock e Papel
		{2, 0}, // Spock vence Tesoura e Pedra
	}

	jogadas := []string{"PEDRA", "PAPEL", "TESOURA", "LAGARTO", "SPOCK"}

	for jogarNovamente == 1 {

		pontosJogador := 0
		pontosPC := 0

		for round := 1; round <= 5; round++ {

			fmt.Println("\n# JOKENPÔ V2 #")
			fmt.Printf("Você: %d | PC: %d\n", pontosJogador, pontosPC)
			fmt.Printf("Round: %d / 5\n\n", round)

			fmt.Println("1 - Pedra")
			fmt.Println("2 - Papel")
			fmt.Println("3 - Tesoura")
			fmt.Println("4 - Lagarto")
			fmt.Println("5 - Spock")
			fmt.Print(">> ")

			var jogador int
			_, err := fmt.Scan(&jogador)
			jogador--

			for err != nil || jogador < 0 || jogador > 4 {
				fmt.Print("Opção inválida! Digite de 1 a 5 >> ")
				_, err = fmt.Scan(&jogador)
				jogador--
			}

			pc := rand.Intn(5)

			fmt.Printf("\nVocê jogou %s e o PC %s.\n", jogadas[jogador], jogadas[pc])

			if jogador == pc {
				fmt.Println("Ninguém ganhou!")
			} else if venceu(vence, jogador, pc) {
				fmt.Println("Você ganhou!")
				pontosJogador++
			} else {
				fmt.Println("O PC ganhou!")
				pontosPC++
			}
		}

		fmt.Println("\n===== PLACAR FINAL =====")
		fmt.Printf("Você: %d | PC: %d\n", pontosJogador, pontosPC)

		fmt.Println("\nJOGAR NOVAMENTE?")
		fmt.Println("1 - Sim")
		fmt.Println("0 - Sair")
		fmt.Print(">> ")

		fmt.Scan(&jogarNovamente)
	}

	fmt.Println("\nObrigado por jogar!")
}

// venceu verifica se a jogada do jogador vence a jogada do pc
func venceu(vence [][]int, jogador, pc int) bool {
	for _, v := range vence[jogador] {
		if v == pc {
			return true
		}
	}
	return false
}