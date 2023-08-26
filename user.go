package todo

type User struct { // Сущность пользователя (поля совпадают с полями в БД)
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
