package qubit

// Qubit é um qubit isolado. Amplitude |0> e |1> ficam aqui depois.
type Qubit struct {
	Index int
}

func New(index int) Qubit {
	return Qubit{Index: index}
}
