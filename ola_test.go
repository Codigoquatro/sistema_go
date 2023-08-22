package main

import "testing"

func Test_Ola(t *testing.T) {
	resultado := Ola("Chris!!")
	esperando := "Olá,Chris!!"

	if resultado != esperando {
		t.Errorf("resultado '%s',esperando '%s'", resultado, esperando)
	}
}
