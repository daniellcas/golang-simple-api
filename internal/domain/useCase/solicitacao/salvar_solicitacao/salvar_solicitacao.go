package salvar_solicitacao

import (
	"errors"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/servico"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/solicitacao"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/status"
)

type SalvarSolicitacaoUseCase struct {
	repositorySolicitacao solicitacao.InterfaceSolicitacaoRepository
	repositoryServico     servico.InterfaceServicoRepository
}

func NovoSalvarSolicitacao(
	solicitacaoRepository solicitacao.InterfaceSolicitacaoRepository,
	servicoRepository servico.InterfaceServicoRepository,
) *SalvarSolicitacaoUseCase {
	return &SalvarSolicitacaoUseCase{
		repositorySolicitacao: solicitacaoRepository,
		repositoryServico:     servicoRepository,
	}
}

func (s *SalvarSolicitacaoUseCase) Execute(input SalvarSolicitacaoInput) (*solicitacao.Solicitacao, error) {
	servico, errServico := s.repositoryServico.PegandoPeloId(input.ServicoId)
	if errServico != nil {
		return nil, errors.New("Erro em encontrar servico")
	}
	status := status.NovoStatus()

	solicitacao := solicitacao.NovaSolicitacao(servico, status, false, input.SolicitanteId)

	errSalvarSolicitacao := s.repositorySolicitacao.Salvar(
		solicitacao.PegandoIdDoServicoDaSolicitacao(),
		solicitacao.PegandoIdDoStatusDaSolicitacao(),
	)

	if errSalvarSolicitacao != nil {
		return nil, errors.New("Erro em salvar solicitacao")
	}

	return solicitacao, nil
}