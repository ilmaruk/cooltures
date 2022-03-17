package names_test

import (
	"github.com/ilmaruk/cooltures/names"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkName(b *testing.B) {
	g := names.New(rand.New(rand.NewSource(time.Now().UnixNano())))
	o := names.Options{
		Gender:  names.GenderMale,
		Culture: names.CultureEnglish,
	}
	for i := 0; i < b.N; i++ {
		name, culture, gender, err := g.Name(o)
		if err != nil {
			panic(err)
		}
		b.Log(name, culture, gender)
	}
}
