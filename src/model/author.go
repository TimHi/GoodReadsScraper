package model

type Author struct {
	AuthorURL     string
	Name          string
	Born          string
	Died          string
	Website       string
	Genre         string
	Description   string
	PhotoURL      string
	AverageRating float64
	Books         []Book
}
