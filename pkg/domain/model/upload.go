package model


type Upload struct {
	ID         string `json:"id"`
	Comment      string `json:"comment"`
	FileURL    string `json:"file_url"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
