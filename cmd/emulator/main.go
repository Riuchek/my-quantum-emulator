package main

import (
	"fmt"
	"my-quantum-emulator/internal/state"
)

func main() {
	fmt.Println("quantum emulator")
	s := state.New(2)
	fmt.Printf("State: %v\n", s.Amplitude)
	s.Hadamard()
	fmt.Printf("State: %v\n", s.Amplitude)
}
