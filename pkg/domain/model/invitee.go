package model


type Invitee struct {
	ID         string `json:"id"`
	FamilyKj      string `json:"family_kj"`
	FirstKj      string `json:"first_kj"`
	FamilyKn  string `json:"family_kn"`
	FirstKn    string `json:"first_kn"`
	Email     string `json:"email"`
	ZipCode     string `json:"zip_code"`
	AddressText    string `json:"address_text"`
	Allergy     string `json:"allergy"`
	UserID     string `json:"user_id"`
	FileURL    string `json:"file_url"`
	UUID       string `json:"uuid"`
	JoinFlag    bool `json:"join_flag"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
