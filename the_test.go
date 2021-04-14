// Run: go test -v

package main

import (
	"testing"

	"github.com/cloudflare/circl/group"
)

func TestScalar(t *testing.T) {
	TestTimes := 100
	t.Log("Test Script")
	G := group.P256
	for i := 0; i < TestTimes; i++ {

		g := G.Generator()
		h := G.Generator()
		if !g.IsEqual(h) {
			t.Fatalf("\nwant: %v\ngot: %v\n", g, h)
		}
	}
}
