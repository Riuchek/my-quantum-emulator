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

func (s *State) Hadamard() {
	h := complex(1/math.Sqrt(2), 0)
	hadamard := [][]complex128{
		{h, h},
		{h, -h},
	}

	// Write into a copy so the second line still sees the original α.
	after := make([]complex128, 2)
	after[0] = hadamard[0][0]*s.Amplitude[0] + hadamard[0][1]*s.Amplitude[1]
	after[1] = hadamard[1][0]*s.Amplitude[0] + hadamard[1][1]*s.Amplitude[1]
	s.Amplitude = after
}
