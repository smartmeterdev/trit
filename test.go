package main

import (
	"fmt"
	"math"
)

func TV() {
	var a, b, c Trit

	tvsum := [27][5]int{{-1, -1, -1, 0, -1},
		{-1, -1, 0, 1, -1},
		{-1, -1, 1, -1, 0},
		{-1, 0, -1, 1, -1},
		{-1, 0, 0, -1, 0},
		{-1, 0, 1, 0, 0},
		{-1, 1, -1, -1, 0},
		{-1, 1, 0, 0, 0},
		{-1, 1, 1, 1, 0},
		{0, -1, -1, 1, -1},
		{0, -1, 0, -1, 0},
		{0, -1, 1, 0, 0},
		{0, 0, -1, -1, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 1, 0},
		{0, 1, -1, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 1, -1, 1},
		{1, -1, -1, -1, 0},
		{1, -1, 0, 0, 0},
		{1, -1, 1, 1, 0},
		{1, 0, -1, 0, 0},
		{1, 0, 0, 1, 0},
		{1, 0, 1, -1, 1},
		{1, 1, -1, 1, 0},
		{1, 1, 0, -1, 1},
		{1, 1, 1, 0, 1}}

	fmt.Println("*********TNOT*********")
	fmt.Println("a\tTnot(a)")
	for a = -1; a <= 1; a++ {
		fmt.Println(a, "\t", TNOT(a))
	}

	fmt.Println("*********TAND*********")
	fmt.Println("a\tb\tand(a,b)")
	for a = -1; a <= 1; a++ {
		for b = -1; b <= 1; b++ {
			fmt.Println(a, "\t", b, "\t", TAND(a, b))
		}
	}

	fmt.Println("*********TOR*********")
	fmt.Println("a\tb\tTOR(a,b)")
	for a = -1; a <= 1; a++ {
		for b = -1; b <= 1; b++ {
			fmt.Println(a, ",", b, ",", TOR(a, b))
		}
	}

	fmt.Println("*********TXOR*********")
	fmt.Println("a,b,S,Ref")
	for a = -1; a <= 1; a++ {
		for b = -1; b <= 1; b++ {
			msg := ""
			if TXOR(a, b) != TXOR_(a, b) {
				msg = "err"
			}
			fmt.Println(a, ",", b, ",", TXOR(a, b), ",", TXOR_(a, b), msg)
		}
	}

	fmt.Println("*********TADD2*********")
	fmt.Println("a,b,Cin,S,Cout")
	i := 0
	for a = -1; a <= 1; a++ {
		for b = -1; b <= 1; b++ {
			for c = -1; c <= 1; c++ {
				s, cout := TADD(a, b, c)
				if int(s) != tvsum[i][3] {
					fmt.Println(a, ",", b, ",", c, ",", s, "=>", tvsum[i][3], cout)
				}
				if int(cout) != tvsum[i][4] {
					fmt.Println(a, ",", b, ",", c, ",", s, ",", cout, "=>", tvsum[i][4])
				}
				i += 1
			}
		}
	}
}

func testSoma() {
	var a, b, c Tint

	a.Set(-6)
	b.Set(6)
	c = TIntAdd(a, b)
	fmt.Println(a.toString(), b.toString())
	fmt.Println(a.ToDecimal(), b.ToDecimal(), c.ToDecimal(), c.toString())
}

func testMaxInt() {

	var a Tint

	a.Set(math.MaxInt64)

	fmt.Println(math.MaxInt64, a.ToDecimal())

	a.Set(math.MinInt64)
	fmt.Println(a.toString())

	fmt.Println(math.MinInt64, a.ToDecimal())
}
func main() {

	//testMaxInt()
	//TV()
	testSoma()

	//fmt.Println(TADD2(-1, -1, 1))
}
