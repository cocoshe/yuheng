package models

import "github.com/google/uuid"

type DisplayRequest struct {
	Idx1 string `json:"idx1"`
	Idx2 string `json:"idx2"`
}

type UserPostDto struct {
	Id   string `json:"id"`
	Post string `json:"post,omitempty"`
	Pic  string `json:"pic"`
}

type CompanyAppealDto struct {
	Id   string `json:"id"`
	Post string `json:"post"`
	Pic  string `json:"pic,omitempty"`
}

type AccusStatusDto struct {
	Id     uuid.UUID `json:"id"`
	Status int       `json:"status"`
}

type AplStatusDto struct {
	Id     uuid.UUID `json:"id"`
	Status int       `json:"status"`
}
