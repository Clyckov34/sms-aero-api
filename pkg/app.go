package pkg

import (
	"fmt"
	"net/http"
)

//Run запуск приложения SMS Aero
//	dictPhone - словарь Образец map[НОМЕР]=ИМЯ
func Run(dictPhone map[string]string) {
	var SMS = &Account{
		Email: "...", // Логин учетной записи
		Token: "...", // Токен учетной записи
		Chan:  make(chan *http.Response),
	}

	for phone, name := range dictPhone {
		go SMS.request(phone, SMS.writeMessage(name))
	}

	fmt.Println(SMS.readStatus())
}
