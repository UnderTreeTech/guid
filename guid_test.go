package guid

import (
	"testing"
)

func TestGuid(t *testing.T) {
	idFactory := NewIDFactory(1)
	idFactory.IdPump()
}

func BenchmarkLoops(b *testing.B) {
	idFactory := NewIDFactory(1)
	for i := 0; i < b.N; i++ {
		idFactory.IdPump()
	}
}

func BenchmarkLoopsParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		idFactory := NewIDFactory(1)
		for pb.Next() {
			idFactory.IdPump()
		}
	})
}
