package names_test

import (
	"github.com/ilmaruk/cooltures/names"
	"testing"
)

func BenchmarkFullNames(b *testing.B) {
	n := names.New()
	o := names.Options{}
	for i := 0; i < b.N; i++ {
		b.Log(n.FullName(o))
	}
}
