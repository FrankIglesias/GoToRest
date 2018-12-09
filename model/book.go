package model

type Book struct {
	ID        int    `json:"id"`
	Author    string `json:"author"`
	Title     string `json:"title"`
	Genre     string `json:"genre"`
	Publisher string `json:"publisher"`
	Year      string `json:"year"`
	ImageURL  string `json:"image_url"`
}

type Books struct {
	Books []Book `json:"books"`
}