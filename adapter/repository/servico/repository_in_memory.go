package servico

import (
	"errors"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/servico"
)

type InMemoryServicoRepository struct {
	Servicos []servico.Servico
}

func (r *InMemoryServicoRepository) Inserir(nome string) (uint, error) {
	_ = append(r.Servicos, servico.Servico{Nome: nome})
	return 1, nil
}

func (r *InMemoryServicoRepository) PegandoPeloId(id uint) (servico.Servico, error) {
	novoSlice := append(r.Servicos, servico.Servico{Id: id, Nome: "TESTE DANIEL"})

	if (int(id)) > len(novoSlice) {
		return servico.Servico{}, errors.New("Servico nao encontrado")
	}
	s := novoSlice[id-1]
	return s, nil
}
