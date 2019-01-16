package utils

import (
	"fmt"
	"math/big"
)

func IntToHex(n *big.Int) string {
	return fmt.Sprintf("0x%x", n)
}
func UIntToHex(n uint64) string {
	return fmt.Sprintf("0x%x", n)
}
