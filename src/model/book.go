package model

type Book struct {
	Authors  []string      `json:"authors"`
	Rating   float64       `json:"rating"`
	Genres   []string      `json:"genres"`
	Title    string        `json:"title"`
	CoverURL string        `json:"coverurl"`
	BookURL  string        `json:"bookurl"`
	Details  EditionDetail `json:"editiondetail"`
}

type EditionDetail struct {
	Format    string `json:"format"`
	Published string `json:"published"`
	ISBN      string `json:"isbn"`
	Language  string `json:"language"`
}
