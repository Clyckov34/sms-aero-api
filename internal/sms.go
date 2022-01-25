package internal

import (
	"fmt"
	"sync"
)

//Run запуск приложения SMS Aero
//	dictPhone - словарь Образец map[НОМЕР]=ИМЯ
//	user - данные пользователя по struct USER
func Run(dictPhone map[string]string, user *User) {
	var SMS = &WaitGroup{
		Account: User{
			Email: user.Email,  // Логин учетной записи
			Token: user.Token}, // Токен учетной записи
		Wg: sync.WaitGroup{},
	}

	SMS.setting(len(dictPhone))

	for phone, name := range dictPhone {
		go SMS.request(phone, writeMessage(name))
	}

	SMS.wait()
}

//writeMessage текст SMS для Абонента
//	 name - Имя абонента
func writeMessage(name string) (message string) {
	return fmt.Sprintf("Привет %v. Меня зовут Клыков Денис, если Вы получили это сообщения, пожалуйста дайте знать об этом. Пишу программу для массовой отправки SMS... Спасибо большое!!!", name)
}
