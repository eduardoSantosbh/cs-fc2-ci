package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(50, 10)
	if total != 20 {
		t.Errorf("Soma deu errado, obtive: %d, esperado: %d.", total, 20)
	}
}
