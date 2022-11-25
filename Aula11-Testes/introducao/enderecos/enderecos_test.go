package enderecos

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "Avenida Paulista"
	tipoDeEnderecoEsperado := "avenida"

	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoEsperado != tipoDeEnderecoRecebido {
		t.Error("O tipo de endereço recebido é diferente do esperado")
	}
}
