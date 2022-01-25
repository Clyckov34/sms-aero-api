package internal

import (
	"sync"
)

type SMSAero interface {
	request(phone, message string)
	setting(number int)
	wait()
}

//Account данные пользователя
type User struct {
	Email string
	Token string
}

//WaitGroup структура для группы горутин
type WaitGroup struct {
	Account User
	Wg      sync.WaitGroup
}
