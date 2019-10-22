// Run: go test -v -run=$^ -bench=. -benchmem

package main

import (
	"crypto/elliptic"
	"testing"
)

func BenchmarkScalarBaseMult(b *testing.B) {
	curve := elliptic.P256()
	k := make([]byte, curve.Params().BitSize/8)
	RandScalar(k)
	for i := 0; i < b.N; i++ {
		_, _ = curve.ScalarBaseMult(k)
	}
}

func BenchmarkScalarMult(b *testing.B) {
	curve := elliptic.P256()
	k := make([]byte, curve.Params().BitSize/8)
	RandScalar(k)
	x0, y0 := curve.ScalarBaseMult(k)
	for i := 0; i < b.N; i++ {
		_, _ = curve.ScalarMult(x0, y0, k)
	}
}
