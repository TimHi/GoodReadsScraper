package model

type Book struct {
	Authors  []string
	Rating   float64
	Genres   []string
	Title    string
	CoverURL string
	BookURL  string
	Pages    int
	Details  EditionDetail
}

type EditionDetail struct {
	Format    string
	Published string
	ISBN      string
	Language  string
}
