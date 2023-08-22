package main

import "testing"

func Test_Ola(t *testing.T) {
	const ola = "Olá"
	resultado := Ola("Chris")
	esperando := ola + "," + "Chris"

	if resultado != esperando {
		t.Errorf("resultado '%s',esperando '%s'", resultado, esperando)
	}
}
