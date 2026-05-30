package drongo

import (
	"fmt"
	"math/rand/v2"
	"reflect"
	"testing"
)

func simpleAddFloat32(a, b, c []float32) {
	// TODO: pre-check so the compiler has a shot at eliding bounds checks.
	for i := range a {
		c[i] = a[i] + b[i]
	}
}

func lessSimpleAddFloat32(a, b, c []float32) {
	if len(a) != len(b) || len(a) != len(c) {
		panic("no good")
	}
	for i := range a {
		c[i] = a[i] + b[i]
	}
}

func randFloat32(n int) []float32 {
	f := make([]float32, n)
	for i := range f {
		f[i] = rand.Float32()
	}
	return f
}

func TestAddFloat32(t *testing.T) {
	var (
		n         = 1001
		a, b      = randFloat32(n), randFloat32(n)
		want, got = make([]float32, n), make([]float32, n)
	)
	simpleAddFloat32(a, b, want)
	simpleAddFloat32(a, b, got)
	// AddFloat32(a, b, got)

	if !reflect.DeepEqual(want, got) {
		t.Fatal("mismatch")
	}
}

func BenchmarkAddFloat32(b *testing.B) {
	for _, size := range []int{16, 65, 259, 1013, 10001} {
		var (
			x, y = randFloat32(size), randFloat32(size)
			z    = make([]float32, size)
		)
		for _, c := range []struct {
			name string
			f    func(a, b, c []float32)
		}{{
			name: "simple",
			f:    simpleAddFloat32,
		}, {
			name: "less_simple",
			f:    lessSimpleAddFloat32,
		}} {
			b.Run(fmt.Sprintf("%05d/%s", size, c.name), func(b *testing.B) {
				for b.Loop() {
					c.f(x, y, z)
				}
			})
		}
	}
}
