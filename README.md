# my-quantum-emulator

Emulador quântico em Go, do zero. Por enquanto é só o esqueleto: qubits, portas, circuito e um pedaço de memória pra brincar com bits.

## Requisitos

- Go 1.27+
- make (opcional)

## Rodar

```bash
make run
```

Equivale a `go run ./cmd/emulator`.

Outros alvos:

| comando      | o que faz                         |
|--------------|-----------------------------------|
| `make build` | gera `bin/emulator`               |
| `make test`  | `go test ./...`                   |
| `make tidy`  | `go mod tidy`                     |
| `make clean` | apaga `bin/`                      |

## Layout

```
cmd/emulator/        ponto de entrada
internal/mem/        bits crus (slice de bytes)
internal/qubit/      um qubit
internal/state/      estado 2^n (só o tamanho por enquanto)
internal/gate/       H, CNOT, X, I — só o nome da porta
internal/circuit/    lista de portas
internal/bell/       circuito H(0) + CNOT(0,1), sem aplicar
```

## Próximo passo

Aplicar as portas no vetor de estado (`internal/state`), depois medir. Bell é H no qubit 0 seguido de CNOT(0, 1).
