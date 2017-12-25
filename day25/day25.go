package day25

type rule struct {
	write     bool
	move      int // 1 right, -1 left
	nextState byte
}

type state struct {
	rules map[bool]rule
}

func (m *machine) extend() {
	if m.position >= 0 {
		if m.position >= len(m.positive) {
			m.positive = append(m.positive, false)
		}
	} else {
		if -m.position-1 >= len(m.negative) {
			m.negative = append(m.negative, false)
		}
	}
}

func (m *machine) read() bool {
	m.extend()
	if m.position >= 0 {
		return m.positive[m.position]
	}
	return m.negative[-m.position-1]
}

func (m *machine) write(b bool) {
	m.extend()
	if m.position >= 0 {
		m.positive[m.position] = b
	} else {
		m.negative[-m.position-1] = b
	}
}

func (m *machine) diagnostic() int {
	count := 0
	for i := 0; i < len(m.positive); i++ {
		if m.positive[i] {
			count++
		}
	}
	for i := 0; i < len(m.negative); i++ {
		if m.negative[i] {
			count++
		}
	}
	return count
}

type machine struct {
	state  byte
	iters  int
	states map[byte]state
	// positive: 0, 1, 2, 3, ...
	// negative: -1, -2, -3, -4, ...
	positive []bool
	negative []bool
	position int
}

func (m *machine) execute() int {
	for i := 0; i < m.iters; i++ {
		rule := m.states[m.state].rules[m.read()]
		m.write(rule.write)
		m.position += rule.move
		m.state = rule.nextState
	}
	return m.diagnostic()
}

func Run() int {
	// Christmas day too hectic to write a parser, just hard code input!
	m := machine{
		'A',
		12861455,
		map[byte]state{
			'A': state{map[bool]rule{
				false: rule{true, 1, 'B'},
				true:  rule{false, -1, 'B'}}},
			'B': state{map[bool]rule{
				false: rule{true, -1, 'C'},
				true:  rule{false, 1, 'E'}}},
			'C': state{map[bool]rule{
				false: rule{true, 1, 'E'},
				true:  rule{false, -1, 'D'}}},
			'D': state{map[bool]rule{
				false: rule{true, -1, 'A'},
				true:  rule{true, -1, 'A'}}},
			'E': state{map[bool]rule{
				false: rule{false, 1, 'A'},
				true:  rule{false, 1, 'F'}}},
			'F': state{map[bool]rule{
				false: rule{true, 1, 'E'},
				true:  rule{true, 1, 'A'}}},
		},
		[]bool{},
		[]bool{},
		0,
	}
	return m.execute()
}
