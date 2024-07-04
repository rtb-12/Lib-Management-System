package types

type BookInfo struct {
	ID                int    `json:"id"`
	Title             string `json:"title"`
	Author            string `json:"author"`
	Genre             string `json:"genre"`
	QuantityAvailable int    `json:"quantityAvailable"`
}

type ListBooks struct {
	Books []BookInfo `json:"books"`
}

type User struct {
	UserID   int    `json:"Userid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogin struct {
	UserID   int    `json:"Userid"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Admin struct {
	AdminId  int    `json:"Adminid"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminLogin struct {
	AdminId  int    `json:"Adminid"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
