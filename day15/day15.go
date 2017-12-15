package day15

type generator struct {
	factor int64
	value  int64
}

func (g *generator) next() {
	g.value = (g.value * g.factor) % 2147483647
}

func score(a, b *generator, iterations int64) int64 {
	s := int64(0)
	for i := int64(0); i < iterations; i++ {
		a.next()
		b.next()
		if a.value%(1<<16) == b.value%(1<<16) {
			s++
		}
	}
	return s
}

func Judge(aValue, bValue, iterations int64) int64 {
	return score(&generator{16807, aValue}, &generator{48271, bValue}, iterations)
}
