package sneakers

type Sneaker struct {
	Id           int      `json:"id"`
	Model        string   `json:"model"`
	Brand        string   `json:"brand"`
	Description  string   `json:"description"`
	Images       []string `json:"images"`
	Color        string   `json:"color"`
	Price        string   `json:"price"`
	Date_release string   `json:"date_Release"`
}

type FavoriteRequest struct {
	Id_user    string `json:"id_user,omitempty"`
	Id_sneaker string `json:"id_sneaker,omitempty"`
}
