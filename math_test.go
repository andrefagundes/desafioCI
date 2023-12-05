package main

import "testing"

func TestSoma(t *testing.T){
	total:= Soma(15,15)
	if total != 30 {
		t.Errorf("Resultado de soma inválido: Resultado %d. Esperado era %d.", total, 30)
	}
}