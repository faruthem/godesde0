package users

import (
	"fmt"
	"time"

	"github.com/farithem/godesde0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Neruda", time.Now(), true)
	fmt.Println("Usuario:", u)
}
