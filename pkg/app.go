package pkg

import (
	"fmt"
	"net/http"
)

//Run запуск приложения SMS Aero
//	dictPhone - словарь Образец map[НОМЕР]=ИМЯ
func Run(dictPhone map[string]string) {
	var SMS = &Account{
		Email: "clyckov.denis2404@gmail.com", // Логин учетной записи
		Token: "lgk2O1lnSBZLPVl83lxnitN7f8B", // Токен учетной записи
		Chan:  make(chan *http.Response),
	}

	for phone, name := range dictPhone {
		go SMS.request(phone, writeMessage(name))
		fmt.Println(SMS.readStatus(), name, phone)
	}
}
