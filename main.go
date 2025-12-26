package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== CAIXA ELETRONICO v1 ===") // isso é falso e não usa dinheiro real, blz?
	fmt.Print("digite seu nome: ")
	var nome string

	// pega o nome completo
	nome, _ = reader.ReadString('\n')
	nome = strings.TrimSpace(nome)

	fmt.Print("digite seu saldo atual: ")
	var saldoAtual float64
	fmt.Scanln(&saldoAtual)

	if saldoAtual < 0 {
		fmt.Println("erro no programa: o valor digitado é invalido")
		return
	}

	// no futuro espero este bloco de codigo vai virar uma função
	fmt.Print("digite o valor que quer sacar: ")
	var valorSaque float64
	fmt.Scanln(&valorSaque)

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