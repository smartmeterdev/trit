package BT

import "log"

type Tint [42]Trit //equivalent int64.  binário 50% a mais de bits

func (t *Tint) Set(n int64) {
	for i := 0; i < len(t); i++ {
		t[i] = 0
	}
	i := len(t) - 1
	for n != 0 && i >= 0 {
		resto := n % base3
		n /= base3

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
		t[i] = Trit(resto)
		i--
	}
	if n != 0 && i < 0 {
		log.Fatal("Over flow")
	}
}

func (t *Tint) Str() (result string) {
	result = ""
	left := false
	for i := 0; i < len(t); i++ {
		if t[i] != 0 {
			left = true
		}
		if left {
			result = result + t[i].Str()
		}
	}
	return
}

func Pow(base, power int64) int64 {
	if power == 0 {
		return 1
	}
	return base * Pow(base, power-1)
}

func (t *Tint) ToDecimal() (result int64) {
	var decimal int64 = 0
	length := len(t)
	// Itera sobre cada caractere do número balanceado
	for i, digit := range t {
		// Soma o valor ponderado ao total
		decimal += int64(digit) * Pow(base3, int64(length-i-1))
	}
	return decimal
}

func TIntAdd(a, b Tint) (s Tint) { // 1848 CMOS - BINÁRIO 2.646 (63 ADD PORT) , binário 43% a mais de CMOS
	var c Trit = 0
	for i := len(a) - 1; i >= 0; i-- {
		s[i], c = TADD(a[i], b[i], c)
	}
	if c != 0 {
		log.Fatal("Overflow")
	}
	return
}
