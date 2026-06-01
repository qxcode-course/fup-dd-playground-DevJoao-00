package main

import "fmt"

type Fator struct {
	num int
	qtd int
}

func calc_fatores(num int) []Fator {
	var fatores []Fator

	div := 2

	for num > 1 {
		cont := 0

		for num % div == 0 {
			cont++
			num /= div
		}

		if cont > 0 {
			fatores = append(fatores, Fator{
				num: div,
				qtd: cont,
			})
		}

		div++
	}

	return fatores
}

func main() {
	var n int

	fmt.Scan(&n)

	fatores := calc_fatores(n)

	for _, f := range fatores {
		fmt.Println(f.num, f.qtd)
	}
}