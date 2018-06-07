package main

import (
	"fmt"
)

func main() {
	var h Hanoi
	h.Run(12)
}

// Hanoi representa a torre de hanoi
type Hanoi struct {
	Rank int
}

// Run realiza a solução da torre de hanoi
func (h *Hanoi) Run(size int) {
	h.solve(size, 1, 3, 2)
}

func (h *Hanoi) solve(size, source, dest, tmp int) {
	if size > 0 {
		h.solve(size-1, source, tmp, dest)
		h.Rank++
		fmt.Printf("%4d ) mover de %d para %d\n", h.Rank, source, dest)
		h.solve(size-1, tmp, dest, source)
	}
}
