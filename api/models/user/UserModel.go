package user

type User struct {
	ID        int    `json:"id"`
	Status    int    `json:"status"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	password  string
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
