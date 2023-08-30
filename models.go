package todo

type User struct { // Сущность пользователя (поля совпадают с полями в БД)
	Id       string `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Segment struct {
	Id   string `json:"-" db:"id"`
	Name string `json:"name" binding:"required"`
}

type UsersSegments struct {
	Id        string
	UserId    string
	SegmentId string
}

type A struct {
	Id     string   `json:"id"`
	Add    []string `json:"add"`
	Delete []string `json:"delete"`
}
