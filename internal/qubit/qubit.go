package qubit

type Qubit struct {
	Index int
}

func New(index int) Qubit {
	return Qubit{Index: index}
}
