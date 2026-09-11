package mem

type Register struct {
	Bits []byte
}

func NewRegister(nbits int) *Register {
	return &Register{
		Bits: make([]byte, nbits),
	}
}

func (r *Register) Len() int {
	return len(r.Bits)
}

func (r *Register) Get(i int) byte {
	return r.Bits[i]
}

func (r *Register) Set(i int, v byte) {
	r.Bits[i] = v
}
