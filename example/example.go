package example

import (
	"fmt"
	"math"
	"math/cmplx"
)

func Example() {
	// 1. Estado Inicial: Qubit está 100% em 0.
	// Index 0 = |0>, Index 1 = |1>
	estadoQubit := []complex128{complex(1, 0), complex(0, 0)}
	fmt.Printf("Estado inicial: |0> = %.2f, |1> = %.2f\n", cmplx.Abs(estadoQubit[0]), cmplx.Abs(estadoQubit[1]))

	// 2. Definir a Porta Hadamard (Matriz 2x2 para criar Superposição)
	invRaizDeDois := complex(1/math.Sqrt(2), 0)
	hadamard := [][]complex128{
		{invRaizDeDois, invRaizDeDois},
		{invRaizDeDois, -invRaizDeDois},
	}

	// 3. Aplicar a Porta (Multiplicação de Matriz x Vetor)
	novoEstado := make([]complex128, 2)
	novoEstado[0] = hadamard[0][0]*estadoQubit[0] + hadamard[0][1]*estadoQubit[1]
	novoEstado[1] = hadamard[1][0]*estadoQubit[0] + hadamard[1][1]*estadoQubit[1]

	// 4. Resultado da Superposição
	fmt.Println("\nApós aplicar a porta Hadamard (Superposição):")

	// Calculando a probabilidade real (|amplitude|^2)
	prob0 := math.Pow(cmplx.Abs(novoEstado[0]), 2)
	prob1 := math.Pow(cmplx.Abs(novoEstado[1]), 2)

	fmt.Printf("Chance de medir 0: %.0f%%\n", prob0*100)
	fmt.Printf("Chance de medir 1: %.0f%%\n", prob1*100)
}
