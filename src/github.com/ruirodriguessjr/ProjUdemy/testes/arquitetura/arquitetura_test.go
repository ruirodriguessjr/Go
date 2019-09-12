package arquitetura

import (
	"runtime"
	"testing"
)

func TesteDependente(t *testing.T) {
	// Faz com que eu execute o teste paralelamente com outra ação tipo uma goroutina(concorrência)
	t.Parallel()
	if runtime.GOARCH == "amd64" {
		t.Skip("Não funciona em arquitetura amd64")
	}
	// ...
	t.Fail()
}
