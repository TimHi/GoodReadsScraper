package model

type Book struct {
	Authors       []string
	Rating        float64
	Genres        []string
	Title         string
	CoverURL      string
	BookURL       string
	Pages         int
	PublishedDate string
	RatingCount   int
	ReviewCount   int
}
