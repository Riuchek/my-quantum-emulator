package gate

type Kind string

const (
	X    Kind = "X"
	H    Kind = "H"
	CNOT Kind = "CNOT"
)

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

func PauliX(q int) Gate {
	return Gate{Kind: X, Targets: []int{q}}
}
