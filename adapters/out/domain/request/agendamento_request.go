package request

type AgendamentoRequest struct {
	DataHoraEnvio string `json:"dataHoraEnvio" binding:"required"`
	Destinatario  string `json:"destinatario" binding:"required"`
	Mensagem      string `json:"mensagem" binding:"required,min=4,max=200"`
}
