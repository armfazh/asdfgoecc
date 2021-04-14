// Run: go test -v -run=$^ -bench=. -benchmem

package main

import (
	"crypto/rand"
	"testing"

	"github.com/cloudflare/circl/group"
)

func BenchmarkScalarBaseMult(b *testing.B) {
	G := group.P256
	g := G.NewElement()
	k := G.RandomScalar(rand.Reader)
	for i := 0; i < b.N; i++ {
		g.MulGen(k)
	}
}

func BenchmarkScalarMult(b *testing.B) {
	G := group.P256
	g := G.Generator()
	k := G.RandomScalar(rand.Reader)
	h := G.NewElement()
	for i := 0; i < b.N; i++ {
		h.Mul(g, k)
	}
}
