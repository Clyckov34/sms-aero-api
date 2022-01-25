package internal

import (
	"fmt"
	"strconv"
	"sync"
)

//Run запуск приложения SMS Aero
//	dictPhone - словарь Образец map[НОМЕР]=ИМЯ
//	user - данные пользователя по struct USER
func Run(listPhone []int, user *User) {
	var SMS = &WaitGroup{
		Account: User{
			Email: user.Email,  // Логин учетной записи
			Token: user.Token}, // Токен учетной записи
		Wg: sync.WaitGroup{},
	}

	SMS.setting(len(listPhone))

	for _, ph := range listPhone {
		go SMS.request(strconv.Itoa(ph), writeMessage())
	}

	SMS.wait()
}

//writeMessage текст SMS для Абонента
//	 name - Имя абонента
func writeMessage() (message string) {
	return fmt.Sprintln("Привет, SMS сообщения от SMS API")
}
