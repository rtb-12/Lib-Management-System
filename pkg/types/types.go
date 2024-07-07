package types

type BookInfo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Stock  int    `json:"stock"`
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
	Password string `password`
}
type AdminLogin struct {
	AdminId  int    `json:"adminid"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
