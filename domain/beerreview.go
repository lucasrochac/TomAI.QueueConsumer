package domain

type Beer struct {
	Name     string  `json:"name"`
	BeerId   int     `json:"beerId"`
	BrewerId int     `json:"brewerId"`
	ABV      float64 `json:"aBV"`
	Style    string  `json:"style"`
}

type Review struct {
	Appearance  float64 `json:"appearance"`
	Aroma       float64 `json:"aroma"`
	Palate      float64 `json:"palate"`
	Taste       float64 `json:"taste"`
	Overall     float64 `json:"overall"`
	Time        int     `json:"time"`
	ProfileName string  `json:"profileName"`
	Text        string  `json:"text"`
}

type BeerReview struct {
	Review Review `json:"review"`
	Beer   Beer   `json:"beer"`
}
