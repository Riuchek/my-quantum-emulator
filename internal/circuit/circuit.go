package circuit

import "my-quantum-emulator/internal/gate"

type Circuit struct {
	NQubits int
	Gates   []gate.Gate
}

func New(nqubits int) *Circuit {
	return &Circuit{NQubits: nqubits}
}

func (c *Circuit) Add(g gate.Gate) {
	c.Gates = append(c.Gates, g)
}
