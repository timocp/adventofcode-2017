package day15

type generator struct {
	factor int64
	value  int64
	picky  int64
}

func (g *generator) next() {
	for {
		g.value = (g.value * g.factor) % 2147483647
		if g.picky == 0 || g.value%g.picky == 0 {
			return
		}
	}
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

func Judge(aValue, bValue, iterations int64, picky bool) int64 {
	a := &generator{16807, aValue, 0}
	b := &generator{48271, bValue, 0}
	if picky {
		a.picky = 4
		b.picky = 8
	}
	return score(a, b, iterations)
}
