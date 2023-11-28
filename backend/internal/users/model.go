package users

type User struct {
	Id         int    `json:"id,omitempty"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	Permission string `json:"permission,omitempty"`
}
