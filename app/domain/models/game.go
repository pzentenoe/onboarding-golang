package models

type Game struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Price    int    `json:"price"`
	ImageURL string `json:"image_url"`
	Year     int    `json:"year"`
}
