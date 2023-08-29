package todo

type User struct { // Сущность пользователя (поля совпадают с полями в БД)
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Segment struct {
	Id   int    `json:"-" db:"id"`
	Name string `json:"name" binding:"required"`
}

type UsersSegments struct {
	Id        int
	UserId    int
	SegmentId int
}

type TodoList struct { // Сущность to-do листа (поля совпадают с полями в БД)
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct { // Сущность элемента to-do листа (поля совпадают с полями в БД)
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}
