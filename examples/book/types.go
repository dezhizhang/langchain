package book

type Section struct {
	Title string `json:"title"`
	Count int    `json:"count"`
}

type BodyCount struct {
	Content     int    `json:"content"`
	LastContent string `json:"lastContent"`
	Count       string `json:"count"`
}
