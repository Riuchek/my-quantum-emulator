package gate

// Kind identifica a porta. Implementação fica pra você.
type Kind string

const (
	I    Kind = "I"
	X    Kind = "X"
	H    Kind = "H"
	CNOT Kind = "CNOT"
)

// Gate aplica uma operação em um ou dois qubits.
type Gate struct {
	Kind    Kind
	Targets []int
}

func Hadamard(q int) Gate {
	return Gate{Kind: H, Targets: []int{q}}
}

func CNot(control, target int) Gate {
	return Gate{Kind: CNOT, Targets: []int{control, target}}
}
