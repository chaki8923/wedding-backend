// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

import (
	"github.com/99designs/gqlgen/graphql"
)

type Mutation struct {
}

type NewInvitation struct {
	Title     string         `json:"title"`
	EventDate string         `json:"event_date"`
	Place     string         `json:"place"`
	Comment   string         `json:"comment"`
	UserID    string         `json:"userId"`
	FileURL   graphql.Upload `json:"file_url"`
}

type NewInvitee struct {
	FamilyKj    string         `json:"family_kj"`
	FirstKj     string         `json:"first_kj"`
	FamilyKn    string         `json:"family_kn"`
	FirstKn     string         `json:"first_kn"`
	Email       string         `json:"email"`
	ZipCode     string         `json:"zip_code"`
	AddressText string         `json:"address_text"`
	Allergy     string         `json:"allergy"`
	UserID      string         `json:"userId"`
	FileURL     graphql.Upload `json:"file_url"`
}

type NewMessage struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type NewUpload struct {
	Comment string         `json:"comment"`
	FileURL graphql.Upload `json:"file_url"`
}

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Query struct {
}

type UpdateInvitation struct {
	ID        string          `json:"id"`
	Title     *string         `json:"title,omitempty"`
	EventDate *string         `json:"event_date,omitempty"`
	Place     *string         `json:"place,omitempty"`
	Comment   *string         `json:"comment,omitempty"`
	FileURL   *graphql.Upload `json:"file_url,omitempty"`
}

type UploadImage struct {
	ID        string `json:"id"`
	Comment   string `json:"comment"`
	FileURL   string `json:"file_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}