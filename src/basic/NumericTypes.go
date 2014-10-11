package main

import (
	"fmt"
	"math/big"
)

const (
	efri          int64 = 10000000000                           // type: int64
	hlutfollum          = 16.0 / 9.0                            // type: float64
	malikvarza          = complex(-2, 3.5) * hlutfollum         // type: complex128
	erGjaldgengur       = 0.0 <= hlutfollum && hlutfollum < 2.0 // type:bool
)

func main() {
	places := handleCommandLine(1000)
	scaledPi := fmt.Sprint(π(places))
	fmt.Printf("3.%s\n", scaledPi[1:])
}

func π(places int) *big.Int {
	digits := big.NewInt(int64(places))
	unity := big.NewInt(0)
	ten := big.NewInt(10)
	exponent := big.NewInt(0)
	unity.Exp(ten, exponent.Add(digits, ten), nil)
	pi := big.NewInt(4)
	left := arccot(big.NewInt(5), unity)
	left.Mul(left, big.NewInt(4))
	right := arccot(big.NewInt(239), unity)
	left.Sub(left, right)
	pi.Mul(pi, left)

	return pi.Div(pi, big.NewInt(0).Exp(ten, ten, nil))
}
