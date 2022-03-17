package names_test

import (
	"github.com/ilmaruk/cooltures/names"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkItalian_Name(b *testing.B) {
	g := names.NewItalian(rand.New(rand.NewSource(time.Now().UnixNano())))
	for i := 0; i < b.N; i++ {
		b.Log(g.Name(names.GenderMale))
	}
}
