package pokeapi

type LocationAreaResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`     // Since this are urls that can be null,
	Previous *string `json:"previous"` // the best way to represent them in GO is with a string pointer
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
