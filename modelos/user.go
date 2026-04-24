package modelos

import (
	"time"
)

type User struct {
	Id        int
	Name      string
	CreatedAt time.Time
	Status    bool
}

func (Usuario *User) AddUser(id int, name string, createAt time.Time, status bool) {
	Usuario.Id = id
	Usuario.Name = name
	Usuario.CreatedAt = createAt
	Usuario.Status = status

}
