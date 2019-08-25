package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) - índices: esperado(%d) <> encontrado (%d)."

// t = referência de objeto para usar minha struct
// testes struct = atributos da minha struct
// a baixo já passando os valores diretamente para os atributos da struct
func TestIndex(t *testing.T) {
	// Faz com que eu execute o teste paralelamente com outra ação tipo uma goroutina(concorrência)
	t.Parallel()
	testes := []struct {
		texto    string // texto passado
		parte    string // parte contendo dentro do texto
		esperado int    // informação da posição para encontrar o se no texto passado parte contem dentro
	}{
		{"Cod3r é show", "Cod3r", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"Rui", "i", 2},
	}

	for _, teste := range testes {
		t.Logf("Massa: %v", teste)
		// atual seria o texto atual passado e
		// se contem parte dentro do texto de acordo com os testes
		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}
