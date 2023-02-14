package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("Kleber")

	esperado := "Olá, Kleber"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}
