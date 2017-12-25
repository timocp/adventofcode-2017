package day25

import "testing"

func TestMachine(t *testing.T) {
	m := machine{
		'A',
		6,
		map[byte]state{
			'A': state{map[bool]rule{
				false: rule{true, 1, 'B'},
				true:  rule{false, -1, 'B'}}},
			'B': state{map[bool]rule{
				false: rule{true, -1, 'A'},
				true:  rule{true, 1, 'A'}}},
		},
		[]bool{},
		[]bool{},
		0,
	}
	r := m.execute()
	if r != 3 {
		t.Errorf("Run() => %d, want 3", r)
	}
}
