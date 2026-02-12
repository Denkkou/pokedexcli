package pokeapi

// This struct is the format the JSON data will be
// unmarshalled into for use in commands
type MapData struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL string `json:"url"`
	} `json:"results"`
}