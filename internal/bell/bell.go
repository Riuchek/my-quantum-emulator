package bell

import (
	"my-quantum-emulator/internal/circuit"
	"my-quantum-emulator/internal/gate"
)

// Bell monta o circuito clássico H + CNOT. Sem simular ainda.
func Circuit() *circuit.Circuit {
	c := circuit.New(2)
	c.Add(gate.Hadamard(0))
	c.Add(gate.CNot(0, 1))
	return c
}
