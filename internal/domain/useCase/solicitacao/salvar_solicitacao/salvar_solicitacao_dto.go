package salvar_solicitacao

type SalvarSolicitacaoInput struct {
	ServicoId uint `json:"servico_id"`
	Campos    []struct {
		Id    uint   `json:"id"`
		Valor string `json:"valor"`
	} `json:"campos"`
	SolicitanteId uint `json:"solicitante_id"`
}

type SalvarSolicitacaoOutput struct {
	Id            uint `json:"id"`
	Concluida     bool `json:"concluida"`
	SolicitanteId uint `json:"solicitante_id"`
	ServicoId     uint `json:"servico_id"`
	Campos        []struct {
		Id    uint   `json:"id"`
		Valor string `json:"valor"`
	} `json:"campos"`
}
