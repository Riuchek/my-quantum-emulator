package main

import (
	"fmt"
	"my-quantum-emulator/internal/state"
)

func main() {
	fmt.Println("quantum emulator")
	s0 := state.New(2)
	fmt.Printf("|00>:     %v\n", s0.Amplitude)
	s0.Hadamard(0)
	fmt.Printf("H no q0:  %v\n", s0.Amplitude)

	s1 := state.New(2)
	s1.Hadamard(1)
	fmt.Printf("H no q1:  %v\n", s1.Amplitude)
}
