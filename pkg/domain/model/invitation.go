package model


type Invitation struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	EventDate  string `json:"event_date"`
	Place      string `json:"place"`
	Comment    string `json:"comment"`
	UserID     string `json:"user_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	FileURL    string `json:"file_url"`
	UUID       string `json:"uuid"`
}
