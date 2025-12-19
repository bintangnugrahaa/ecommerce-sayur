package response

type ListResponse struct {
	ID      uint   `json:"id"`
	Subject string `json:"subject"`
	Status  string `json:"status"`
	SentAt  string `json:"sent_at"`
}
