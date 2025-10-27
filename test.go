package main

import (
	"fmt"
	"math"
	"trit/BT"
)

func TV() {
	var a, b, c BT.Trit

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
		fmt.Println(a, "\t", BT.TNOT(a))
	}

	fmt.Println("*********TAND*********")
	fmt.Println("a\tb\tand(a,b)")
	for a = -1; a <= 1; a++ {
		for b = -1; b <= 1; b++ {
			fmt.Println(a, "\t", b, "\t", BT.TAND(a, b))
		}
	}

	fmt.Println("*********TOR*********")
	fmt.Println("a\tb\ttor(a,b)")
	for a = -1; a <= 1; a++ {
		for b = -1; b <= 1; b++ {
			fmt.Println(a, "\t", b, "\t", BT.TOR(a, b))
		}
	}

	fmt.Println("*********TXOR*********")
	fmt.Println("a\tb\ttxor(a,b)")
	for a = -1; a <= 1; a++ {
		for b = -1; b <= 1; b++ {
			fmt.Println(a, "\t", b, "\t", BT.TXOR(a, b))
		}
	}

	fmt.Println("*********TADD*********")
	fmt.Println("a,b,Cin,S,Cout")
	i := 0
	for a = -1; a <= 1; a++ {
		for b = -1; b <= 1; b++ {
			for c = -1; c <= 1; c++ {
				s, cout := BT.TADD(a, b, c)
				if int(s) != tvsum[i][3] {
					fmt.Println(a, ",", b, ",", c, ",", s, "=>", tvsum[i][3], cout)
				}
				if int(cout) != tvsum[i][4] {
					fmt.Println(a, ",", b, ",", c, ",", s, ",", cout, "=>", tvsum[i][4])
				}
				fmt.Println(a, ",", b, ",", c, ",", s, ",", cout)
				i += 1
			}
		}
	}
}

func testSoma() {
	var a, b, c BT.Tint

	a.Set(-6)
	b.Set(6)
	c = BT.TIntAdd(a, b)
	fmt.Println(a.Str(), b.Str())
	fmt.Println(a.ToDecimal(), b.ToDecimal(), c.ToDecimal(), c.Str())
}

func testMaxInt() {

	var a BT.Tint

	a.Set(math.MaxInt64)

	fmt.Println(math.MaxInt64, a.ToDecimal())

	a.Set(math.MinInt64)
	fmt.Println(a.Str())

	fmt.Println(math.MinInt64, a.ToDecimal())
}
func main() {

	//testMaxInt()
	TV()
	//testSoma()

	//fmt.Println(TADD2(-1, -1, 1))
}
