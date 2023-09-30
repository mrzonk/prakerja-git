package cerpen

type Cerpen struct {
	Id    int    `json:"id"`
	title string `json:"judul"`
	body  string `json:"body"`
}

type ResponseCerpen struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
