package model

type Movie struct {
	Id          string
	Title       string
	Image       string
	Description string
}

func NewMovie() *Movie {
	return &Movie{}
}
