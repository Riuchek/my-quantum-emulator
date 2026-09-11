# Lição 1 — o vetor de estado 

Esta lição cobre o que o `New` em `internal/state` faz hoje.
O `example/example.go` é só um rascunho de matemática; o estado de verdade mora neste vetor.

## O que você está representando

Um qubit não é `true`/`false`. Ele é

\[
|\psi\rangle = \alpha|0\rangle + \beta|1\rangle
\]

\(\alpha\) e \(\beta\) são **números complexos** (amplitudes). A probabilidade de medir 0 é \(|\alpha|^2\); de medir 1 é \(|\beta|^2\). As duas somam 1.

Isso é um **vetor** (uma lista de amplitudes), não uma matriz. Matriz é porta (Hadamard, X, …): ela *multiplica* o vetor. Ainda não estamos nisso.

No print do Go, um qubit em \(|0\rangle\) aparece assim:

```
[(1+0i) (0+0i)]
```

Leia como coluna:

\[
|\psi\rangle = \begin{pmatrix} \alpha \\ \beta \end{pmatrix}
= \begin{pmatrix} 1 \\ 0 \end{pmatrix}
= |0\rangle
\]

- índice `0` → \(\alpha\) → amplitude de \(|0\rangle\) (ou \(|00\ldots0\rangle\) se houver mais qubits)
- índice `1` → \(\beta\) → amplitude de \(|1\rangle\)

`complex(real, imag)` monta `real + imag·i`. Não é o par “(chance de 0, chance de 1)”.
`complex(1, 0)` é o número **1**. `complex(0, 1)` é o número **i** — não é o estado \(|1\rangle\).

\(|1\rangle\) puro seria `[(0+0i) (1+0i)]`, com `complex(1, 0)` no **segundo** índice.

## `|0⟩` não é superposição

A fórmula \(\alpha|0\rangle + \beta|1\rangle\) vale sempre. Superposição é quando **os dois** são diferentes de zero.

| estado | \(\alpha\) | \(\beta\) | print |
|--------|------------|-----------|--------|
| \(|0\rangle\) puro (inicial) | 1 | 0 | `[(1+0i) (0+0i)]` |
| \(|1\rangle\) puro | 0 | 1 | `[(0+0i) (1+0i)]` |
| superposição 50/50 (depois de H) | \(1/\sqrt{2}\) | \(1/\sqrt{2}\) | ~ `[(0.71+0i) (0.71+0i)]` |

\(\alpha = 1\), \(\beta = 0\) = certeza de medir 0. Não misture com “ket alpha = superposição”.

## Por que o slice tem tamanho \(2^n\), não \(n\)

Cada qubit **dobra** o número de estados da base.

| qubits | estados da base | tamanho do slice |
|--------|-----------------|------------------|
| 1 | \(|0\rangle\), \(|1\rangle\) | 2 |
| 2 | \(|00\rangle\), \(|01\rangle\), \(|10\rangle\), \(|11\rangle\) | 4 |
| 3 | oito kets | 8 |

Dois qubits **não** são dois complexos. São **um** vetor de 4 amplitudes. Emaranhamento não cabe em variáveis `qubitA` / `qubitB` separadas.

No código isso é:

```go
numEstadosDaBase := 1 << numQubits
```

`1 << numQubits` é conta de **inteiro em Go** (empurrar o bit do número `1` para a esquerda `n` vezes = multiplicar por \(2^n\)). Não é “mexer no bit na memória do qubit”. O registrador quântico ainda nem existe nessa linha: você só descobre **quantas gavetas** o vetor precisa.

Nomes tipo `dim` escondem isso. `dim` em álgebra linear *é* a dimensão do espaço (\(2^n\)), mas no código vale o nome que você lê em voz alta: **número de estados da base**.

## O que o `make` faz

```go
amplitudes := make([]complex128, numEstadosDaBase)
```

Isso só **aloca** a lista: `numEstadosDaBase` casas de `complex128`. O zero de complexo em Go é `0+0i`. Com 1 qubit, depois do `make`:

```
índice:     0         1
valor:    0+0i      0+0i
```

Ainda não é um estado físico útil. É gaveta vazia do tamanho certo.

```go
amplitudes[0] = complex(1, 0)
```

Só a primeira casa vira 1. O resto **já era zero**; você não “seta tudo zero = 1”. Fica:

```
índice:     0         1
valor:    1+0i      0+0i
```

O tamanho **trava** em \(2^n\). Portas depois **reescrevem** esses números; não nascem índices novos no meio do circuito.

Se você chumbasse `amplitudes[1]` no `New`:

- com 1 qubit, arriscaria colocar \(|1\rangle\) ou até `i` no lugar errado;
- com 2 qubits, o índice `1` é a amplitude de \(|01\rangle\), não “o segundo qubit”;
- com 0 qubits, `amplitudes[1]` dá panic.

Estado inicial padrão: **só** índice 0 vale 1 → \(|00\ldots0\rangle\).

## Conferência do passo 1

`New(1)` deve imprimir `[(1+0i) (0+0i)]`.

Isso fecha o vetor inicial. Superposição entra no passo 2: aplicar Hadamard **nesse** vetor (ainda em `state`, ainda sem `gate`/`circuit`/`mem`).
