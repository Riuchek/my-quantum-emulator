package state

// State é o estado do sistema (vetor de amplitudes, 2^n entradas).
type State struct {
	NQubits int
}

func New(nqubits int) *State {
	return &State{NQubits: nqubits}
}

func (s *State) Dim() int {
	return 1 << s.NQubits
}
