// Run: go test -v

package main

import (
	"crypto/elliptic"
	"testing"
)

func TestScalar(t *testing.T) {
	TestTimes := 100
	t.Log("Test Script")

	curve := elliptic.P256()
	x0, y0 := curve.Params().Gx, curve.Params().Gy
	k := make([]byte, curve.Params().BitSize/8)

	for i := 0; i < TestTimes; i++ {
		RandScalar(k)
		x1, y1 := curve.ScalarBaseMult(k)
		x2, y2 := curve.ScalarMult(x0, y0, k)
		if x1.Cmp(x2) != 0 || y1.Cmp(y2) != 0 {
			t.Fatalf("\nwant: %v\ngot: %v\n", HexPoint(x1, y1), HexPoint(x2, y2))
		}
	}
}
