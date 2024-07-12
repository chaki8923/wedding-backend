package model

type SendMail struct {
	ID        string `json:"id"`
	To      string `json:"to"`
	From    string `json:"from"`
	Subject    string `json:"subject"`
	Body    string `json:"body"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
