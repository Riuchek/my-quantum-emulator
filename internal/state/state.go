package state

import "math"

// State is the state vector of the entire system (not a single qubit).
// Amplitudes[i] is the complex amplitude of the state of the base with index i.
// With 1 qubit: [α, β] = amplitudes of |0⟩ and |1⟩.
// With n qubits the slice has 2^n entries (one number per ket of the base).
type State struct {
	NumQubits int
	Amplitude []complex128
}

func New(numQubits int) *State {
	// 1 << n = 2^n = how many kets the base has (|00…0⟩, |00…1⟩, …).
	numStatesOfBase := 1 << numQubits
	amplitude := make([]complex128, numStatesOfBase)
	// |00…0⟩: only the first amplitude is 1; the make already zeroed the rest.
	amplitude[0] = complex(1, 0)
	return &State{NumQubits: numQubits, Amplitude: amplitude}
}

// Hadamard applies H to one qubit. Qubit 0 is the least significant bit of the index:
// index 0 = |…00⟩, 1 = |…01⟩, 2 = |…10⟩, 3 = |…11⟩.
func (s *State) Hadamard(qubit int) {
	h := complex(1/math.Sqrt(2), 0)
	after := make([]complex128, len(s.Amplitude))
	bitOfQubit := 1 << qubit

	for i := range s.Amplitude {
		if i&bitOfQubit != 0 {
			continue
		}
		partner := i | bitOfQubit
		a := s.Amplitude[i]
		b := s.Amplitude[partner]
		after[i] = h*a + h*b
		after[partner] = h*a - h*b
	}
	s.Amplitude = after
}
