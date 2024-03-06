package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Rsultado invalido: Resultado %d. Esperadp %d", total, 30)
	}
}
