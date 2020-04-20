package main

import "testing"

func TestSoma(t *testing.T) {
    resultado := Sum(5, 5)
    if resultado != 10 {
       t.Errorf("Soma incorreta, Resultado: %d, Esperado: %d.", resultado, 10)
    }
}