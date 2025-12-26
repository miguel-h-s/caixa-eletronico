package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// variavel global
var valorSaque float64

func digiteSaqueValor() {
	fmt.Print("digite o valor que quer sacar: ")
	fmt.Scanln(&valorSaque)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("=== CAIXA ELETRONICO v2 ===") // isso é falso e não usa dinheiro real, blz?

		fmt.Println(" '/exit' para sair ")

		fmt.Print("digite seu nome: ")
		var nome string

		// pega o nome completo
		nome, _ = reader.ReadString('\n')
		nome = strings.TrimSpace(nome)

		if nome == "/exit" {
			break
		}

		fmt.Print("digite seu saldo atual: ")
		var saldoAtual float64
		fmt.Scanln(&saldoAtual)

		if saldoAtual < 0 {
			fmt.Println("erro no programa: o valor digitado é invalido")
			return
		}

		digiteSaqueValor()

		// validação de entrada
		if valorSaque <= 0 {
			fmt.Println("erro no programa: o valor digitado é invalido")
			return
		} else if valorSaque > saldoAtual {
			fmt.Println("erro no programa: saque digitado maior que seu saldo atual")
			return
		} else {
			fmt.Println("saldo autorizado")
			saldoAtual -= valorSaque
			fmt.Printf("saldo atual:  %.2f\n", saldoAtual)
			fmt.Printf("valor do saque: %.2f\n", valorSaque)
		}
	}
}
