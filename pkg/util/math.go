package util

import (
	"log"
	"math/big"
)

func Pow10(n int) *big.Int {
	return big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(n)), nil)
}

func DivideBigInt(num1 *big.Int, num2 *big.Int) *big.Float {
	if num2.BitLen() == 0 {
		log.Fatal("cannot divide by zero.")
	}
	num1BigFloat := new(big.Float).SetInt(num1)
	num2BigFloat := new(big.Float).SetInt(num2)
	result := new(big.Float).Quo(num1BigFloat, num2BigFloat)
	return result
}
