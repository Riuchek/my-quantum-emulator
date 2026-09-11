package bell

import (
	"my-quantum-emulator/internal/circuit"
	"my-quantum-emulator/internal/gate"
)

func BellPairCircuit() *circuit.Circuit {
	c := circuit.New(2)
	c.Add(gate.Hadamard(0))
	c.Add(gate.CNot(0, 1))
	return c
}
