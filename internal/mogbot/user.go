package mogbot

type User struct {
	ID string
}

type UserService interface {
	AddUser(*User) error
	GetUser(string) (*User, error)
	GetAllUsers() ([]*User, error)
	UpdateUser(*User) error
	RemoveUser(string) error
}
