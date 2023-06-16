package main

import (
	"fmt"
	"math/rand"
	"time"
)

func gerarnum() int {
	return rand.Intn(100) + 1
}

func chutes() int {
	var Chute int
	_, err := fmt.Scanln(&Chute)
	for err != nil {
		fmt.Println("Por favor, digite um número válido.")
		_, err = fmt.Scanln(&Chute)
	}
	return Chute
}

var jogosnum [][]int

func registrarJogada(tentativas []int) {
	jogosnum = append(jogosnum, tentativas)
}

func exibirJogadas() {
	for i, tentativas := range jogosnum {
		fmt.Printf("Jogada %d - Tentativas: %v\n", i+1, tentativas)
	}
}
func Jogo() {
	resposta := gerarnum()
	tentativas := []int{}
	numeroTentativas := 0

	fmt.Println("Tente adivinhar o número (entre 1 e 100):")

	for {
		Chute := chutes()

		numeroTentativas++
		tentativas = append(tentativas, Chute)

		if Chute == resposta {
			fmt.Printf("Parabéns! Você acertou o número em %d tentativas.\n", numeroTentativas)
			registrarJogada(tentativas)
			break
		} else if Chute < resposta {
			fmt.Println("O número é maior.")
		} else {
			fmt.Println("O número é menor.")
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Bem-vindo ao jogo de adivinhação!")

	for {
		Jogo()

		fmt.Println("Deseja jogar novamente? (s/n)")
		var escolha string
		fmt.Scanln(&escolha)

		if escolha != "s" && escolha != "S" {
			break
		}
	}

	fmt.Println("\nEstatísticas das jogadas:")
	exibirJogadas()
}
