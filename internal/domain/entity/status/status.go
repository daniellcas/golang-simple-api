package status

const (
	SolicitacaoIniciada    uint = 1
	ESPERA_DA_APROVACAO         = 2
	SOLICITACAO_FINALIZADA      = 3
)

type Status struct {
	Id   uint
	nome string
}

func NovoStatus() Status {
	return Status{
		Id:   SolicitacaoIniciada,
		nome: "Solicitacao enviada",
	}
}
