package main

// Balanced Ternary Representation
// Trit type to represent a single trit (-1, 0, 1)
type Trit int8

const AnalogSlice float64 = 0.078 //FOR -5V .. 5V INPUT
const base3 = 3

// Logical ports
func TNOT(a Trit) Trit { //2 CMOS
	return -a
}

func TNOR(a, b Trit) Trit { //4 CMOS - BINÁRIO 6
	if a == 1 || b == 1 {
		return -1
	} else if a == -1 || b == -1 {
		return 1
	}
	return 0
}

func TOR(a, b Trit) Trit { // 6 CMOS - BINÁRIO 6
	return TNOT(TNOR(a, b))
}

func TNAND(a, b Trit) Trit { // 4 CMOS - BINÁRIO 6
	if a == 1 && b == 1 {
		return -1
	} else if a == -1 && b == -1 {
		return 1
	}
	return 0
}

func TAND(a, b Trit) Trit { // 6 CMOS - BINÁRIO 6
	return TNOT(TNAND(a, b))
}

func TXOR_(a, b Trit) Trit { //simulado
	base3 := Trit((a.NBal() + b.NBal()) % base3)
	if base3 == 2 {
		return -1
	}
	return base3
}

func TXOR(a, b Trit) Trit { /// 10 CMOS 3 TRIAC - BINÁRIO 12
	s1 := TNAND(a, b)
	if a != 0 && b != 0 {
		return s1
	}
	if s1 == 0 {
		return TOR(a, b)
	}
	return 0
}

func TABS(a Trit) Trit { //2 CMOS
	if a == 1 || a == -1 {
		return 1
	}
	return 0
}

func TADD(a, b, cin Trit) (s Trit, cout Trit) { // 44 CMOS - BINÁRIO 42
	s1 := TXOR(a, b)
	s = TXOR(s1, cin)
	s2 := TAND(s1, cin)
	s3 := TAND(a, b)
	cout = TXOR(s2, s3)
	//fmt.Println(s2, s3, cout)
	return
}

func (t *Trit) toString() string {
	switch *t {
	case -1:
		return "-"
	case 1:
		return "+"
	default:
		return "0"
	}
}

func (t *Trit) Analog() float64 {
	return float64(t.NBal()) * AnalogSlice
}

func (t *Trit) NBal() Trit {
	if *t == -1 {
		return 2
	}
	return *t
}
