// Run go run main.go

package main

import (
	"crypto/elliptic"
	"fmt"
	"math/big"
	"math/rand" // "crypto/rand"
)

func HexBytes(x []byte) string {
	n := len(x)
	s := ""
	for i := range x {
		s += fmt.Sprintf("%02x", x[n-i-1])
	}
	return s
}

func HexPoint(x, y *big.Int) string {
	return fmt.Sprintf("\nx: %v\ny: %v", x.Text(16), y.Text(16))
}

func RandScalar(k []byte) { rand.Read(k) }

func main() {

	fmt.Print("Scalar Multiplication\n")

	curve := elliptic.P256()

	k := make([]byte, curve.Params().BitSize/8)
	RandScalar(k)
	fmt.Printf("k: %v\n", HexBytes(k))

	x, y := curve.ScalarBaseMult(k)
	fmt.Printf("P: %v\n", HexPoint(x, y))
}
