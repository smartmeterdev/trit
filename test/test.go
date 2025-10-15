package main

import (
	"fmt"
	"math"
	"strings"
)

// Função para converter um número decimal para ternário balanceado
func decimalParaTernarioBalanceado(n int) string {
	if n == 0 {
		return "0"
	}

	var resultado []string

	for n != 0 {
		resto := n % 3
		n /= 3

		// Ajusta o resto para o sistema ternário balanceado
		switch resto {
		case 2:
			resto = -1
			n++ // Compensa o ajuste
		case -2:
			resto = 1
			n-- // Compensa o ajuste
		}

		// Adiciona o dígito correspondente ao resultado
		if resto == -1 {
			resultado = append(resultado, "T") // Representa -1 como 'T'
		} else {
			resultado = append(resultado, fmt.Sprintf("%d", resto))
		}
	}

	// Inverte o resultado para obter a ordem correta
	for i, j := 0, len(resultado)-1; i < j; i, j = i+1, j-1 {
		resultado[i], resultado[j] = resultado[j], resultado[i]
	}

	return strings.Join(resultado, "")
}

func main() {
	var numero int

	// Solicita ao usuário um número decimal
	fmt.Print("Digite um número decimal: ")
	_, err := fmt.Scan(&numero)
	if err != nil {
		fmt.Println("Erro ao ler o número. Certifique-se de digitar um número inteiro.")
		return
	}

	// Converte para ternário balanceado e exibe o resultado
	ternarioBalanceado := decimalParaTernarioBalanceado(numero)
	fmt.Printf("O número %d em ternário balanceado é: %s\n", numero, ternarioBalanceado)

	// Converte para decimal
	decimal, err := balancedTernaryToDecimal(ternarioBalanceado)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	// Exibe o resultado
	fmt.Printf("O número decimal equivalente é: %d\n", decimal)

}

// balancedTernaryToDecimal converte um número em ternário balanceado para decimal
func balancedTernaryToDecimal(input string) (int, error) {
	// Verifica se a entrada é válida
	for _, char := range input {
		if char != 'T' && char != '0' && char != '1' {
			return 0, fmt.Errorf("entrada inválida: '%s'. Use apenas '-', '0' e '1'", input)
		}
	}

	decimal := 0
	length := len(input)

	// Itera sobre cada caractere do número balanceado
	for i, char := range input {
		// Calcula a potência correspondente
		power := length - i - 1

		// Determina o valor do dígito
		var digit int
		switch char {
		case '1':
			digit = 1
		case '0':
			digit = 0
		case 'T':
			digit = -1
		}

		// Soma o valor ponderado ao total
		decimal += digit * int(math.Pow(3, float64(power)))
	}

	return decimal, nil
}
