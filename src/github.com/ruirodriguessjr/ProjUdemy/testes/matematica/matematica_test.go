package matematica

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

// Métodos para fazer teste começam com Test e o arquivo tem que ter _teste.go
func TestMedia(t *testing.T) {
	// Faz com que eu execute o teste paralelamente com outra ação tipo uma goroutina(concorrência)
	t.Parallel()
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}

}
