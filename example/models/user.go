package models

type User struct {
	Id   int64 `json:"id"`
	Name string
}

func (u *User) EntityName() string {
	return "user"
}

func (u *User) ID() any {
	return u.Id
}
